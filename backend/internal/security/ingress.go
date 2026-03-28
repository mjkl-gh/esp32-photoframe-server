package security

import (
	"net"
	"os"

	"github.com/labstack/echo/v4"
)

// IsTrustedIngressRequest returns true when a request is coming through
// Home Assistant ingress from the trusted Supervisor ingress proxy.
func IsTrustedIngressRequest(c echo.Context) bool {
	if c.Request().Header.Get("X-Ingress-Path") == "" {
		return false
	}

	trustedProxyIP := os.Getenv("HA_INGRESS_PROXY_IP")
	if trustedProxyIP == "" {
		trustedProxyIP = "172.30.32.2"
	}

	trustedIP := net.ParseIP(trustedProxyIP)
	if trustedIP == nil {
		return false
	}

	remoteHost, _, err := net.SplitHostPort(c.Request().RemoteAddr)
	if err != nil {
		remoteHost = c.Request().RemoteAddr
	}

	remoteIP := net.ParseIP(remoteHost)
	return remoteIP != nil && remoteIP.Equal(trustedIP)
}
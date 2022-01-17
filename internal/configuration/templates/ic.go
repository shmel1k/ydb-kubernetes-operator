package templates

const InterconnectConfigTemplate = `
StartTcp: true
{{- if .Spec.Service.Interconnect.TLSConfiguration.Enabled }}
EncryptionMode: REQUIRED
{{- end }}
MaxInflightAmountOfDataInKB: 10240
HandshakeTimeoutDuration {
  Seconds: 1
}
{{- if .Spec.Service.Interconnect.TLSConfiguration.Enabled }}
PathToCertificateFile: "/tls/interconnect/tls.crt"
PathToPrivateKeyFile: "/tls/interconnect/tls.key"
PathToCaFile: "/tls/interconnect/ca.crt"
{{- end }}
`

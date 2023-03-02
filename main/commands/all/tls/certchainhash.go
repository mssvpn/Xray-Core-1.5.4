package tls

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/mssvpn/Xray-Core-1.5.4/main/commands/base"
	"github.com/mssvpn/Xray-Core-1.5.4/transport/internet/tls"
)

var cmdCertChainHash = &base.Command{
	UsageLine: "{{.Exec}} certChainHash",
	Short:     "Calculate TLS certificates hash.",
	Long: `
	xray tls certChainHash --cert <cert.pem>
	Calculate TLS certificate chain hash.
	`,
}

func init() {
	cmdCertChainHash.Run = executeCertChainHash // break init loop
}

var input = cmdCertChainHash.Flag.String("cert", "fullchain.pem", "The file path of the certificates chain")

func executeCertChainHash(cmd *base.Command, args []string) {
	fs := flag.NewFlagSet("certChainHash", flag.ContinueOnError)
	if err := fs.Parse(args); err != nil {
		fmt.Println(err)
		return
	}
	certContent, err := ioutil.ReadFile(*input)
	if err != nil {
		fmt.Println(err)
		return
	}
	certChainHashB64 := tls.CalculatePEMCertChainSHA256Hash(certContent)
	fmt.Println(certChainHashB64)
	return
}

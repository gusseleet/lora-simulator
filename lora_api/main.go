package main


import (
	"context"
	"log"
	"github.com/brocaar/lora-app-server/api"
	"google.golang.org/grpc"
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/metadata"
)


type Authentication struct {
	authorization string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"authorization":    a.authorization,
	}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func main() {

	ctx := context.Background()
	info, err := getTLS()
	if err != nil {
		log.Fatal(err)
	}

	auth := Authentication{
		authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJhdWQiOiJsb3JhLWFwcC1zZXJ2ZXIiLCJuYmYiOjE1MTkwNDM3MjIsImV4cCI6MTUxOTEzMDEyMiwic3ViIjoidXNlciIsInVzZXJuYW1lIjoiYWRtaW4ifQ.s5nptwaPyeAQJncdbu-7pWKHxNF2IReWPYd_56MJIlQ",
	}


	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(info)),
		grpc.WithPerRPCCredentials(&auth),
	}

	conn, err := grpc.Dial("18.194.141.23:8080", grpcOpts...)
	if err != nil {
		fmt.Println("Error in connect")
		log.Fatal(err)
	}

	t := api.NewUserClient(conn)
	k, err := t.List(ctx, &api.ListUserRequest{Limit:100, Offset:0})

	if err != nil {
		fmt.Println(err)
	}


	for _, element := range k.Result {
		fmt.Println(element)
	}


	//a := api.NewApplicationClient(conn)
/*
	context.WithValue()

	md, ok := metadata.FromIncomingContext(context.Background())



	if !ok {
		fmt.Print("Error in meta")
	}

	token, ok := md["authorization"]

	fmt.Println(token)

	r, err := t.Create(ctx, &api.AddUserRequest{Username:"testGRPC", Password:"myPassword", IsAdmin:false, Email:"testuser@gmail.com"})
*/

/*

	if err != nil {
		fmt.Println("Error in create")
		log.Fatal(err)
	}
*/

	/*response, err := a.Create(ctx, &api.CreateApplicationRequest{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(response)
*/

}

func getTLS() (*tls.Config, error) {

	certPool := x509.NewCertPool()


	t := `-----BEGIN CERTIFICATE-----
MIIEpDCCAowCCQDo1vuFenkFzDANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDDAls
b2NhbGhvc3QwHhcNMTgwMjEzMTUwNjEyWhcNMTkwMjEzMTUwNjEyWjAUMRIwEAYD
VQQDDAlsb2NhbGhvc3QwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQC5
VUQUHGAQteApiU9tlhVdBD+nVCfA+w9SRoEYEYGuOHzYNhXzBhWUJcCRz9hL9wLZ
27wyYAPNgCx6H3hu/MzcKA4VFf7YEPaAY4J1jeK+CzKkLpyk6N4r6NN3ubWsjEHn
sMrtwLteUMfjUi4vtwJrXt9f4Nl50IZt13nd+1FmqZfA2E+PlsP62MfMBjjjcVZ1
rXui1nOVrSh9oGD6xoIifQ6EPxBE2Ztt1OCTVs/Ab0nKJjHZiFZRGCws4t3GIla/
hz/d2V0FIBDkhQLaqM2Ti0Fpopw8JrTck0x9e141tFnPkiLbDvqH8bxPwdVUGmC7
SynzUfmTyE60ENu4rI7aHgiFdEzhfcOtAc9vfIR3eKUIM23l7dZ84qHbTsx30X/I
nHWgJLWCPxQ8F9yeev4m/ducWU8i9aU9GZ9XIr7u5bYfKBKHGg8nZMO/OgXxzW5Q
qZRzB+WMPHLY20rv/J6FNYdvuWKLjPcqEk66/mXBKKEGRxjPay1NUAHNM1kaIY6c
xQk2tOaH4b5H7Vbb6MIrQGaZI+2fvzFPkHtmSyUeWmqw8B9N7aoCf0z/3IqJrhnY
JYfgWD9j5WpIGwH8EEfY1Ekza+YkOOOMkm1hrIacwVSvRy4vICPFRAx2axWIsKO3
hl21NwpVX0LEcRLS2res6Zx4dc8oyfe31+Fj3rm8hQIDAQABMA0GCSqGSIb3DQEB
CwUAA4ICAQAv3FqD850V+joVDJRCYFXlYn0sdD/D3vmg+uLwHpM72x/jjRQCMtnB
qjgbPltrsPQ59UoRTKvgNiNFXxXI3hIJnUe6jFCD+ZVE1sFoJFh4r4VdHyxpxjkD
zClRTiZFI0+E53kPm5PmIaq9i6+qVoxKmciX87C3I8mwEc/szXdNlOfGSzFhbPzy
6hT0L2tLvCHu3am2H/vmRo17DOGPhD1xdKmH3LoeROxoai8FvInfocFtDqzj/KcM
vCM8VwIP4vrM2WIcL9LvCK8hwvGlkdlaFgJeX5fKyjWonesmgfuhgonfZ327ZqND
3EIcaH5rXwBFWe9rvGxtaLAtIYEfeSd8gxMzXplcN3ETmOOpyvrNQScJUQuSXFnt
g3FyuFl6Kjjihj4ksYyqfJocLwvvnZWCdOSBpkBKK5KD1ENoE0FtC62+fLLOwvPJ
pm93nrOZMvQ0WoZP5UxGqhC662vvYQuhjsGCScv5wJYx2PvOWNl8dbpFcaKW+sof
OMT6AJhE/KTflON1IzftK+woXPuM7iVWBywZxCl2H5C57NHaQRtw1Jc0LbKFuTnP
1cz5TKDwmRBAIw/jUowsbMrjTdjkfU9gGPhr1Q02Me97hnnDMRD86CKO3Og7Go1Q
TKqoxGJsKP1aDNDPFvV099LZMSwzJBYcmveLAPOabGYQbNlOgd44xQ==
-----END CERTIFICATE-----`


	if ok := certPool.AppendCertsFromPEM([]byte(t)); !ok {
		log.Fatal("Error", ok)
		return nil, nil
	}


	return &tls.Config{ServerName: "localhost", RootCAs: certPool}, nil
}



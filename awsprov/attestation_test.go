package awsprov

import (
	"encoding/base64"
	"strings"
	"testing"
)

const EC2_METADATA_AMI = "ami-ea165990"
const EC2_RSA2048 = "MIAGCSqGSIb3DQEHAqCAMIACAQExDzANBglghkgBZQMEAgEFADCABgkqhkiG9w0BBwGggCSABIIB1HsKICAiYXZhaWxhYmlsaXR5Wm9uZSIgOiAidXMtZWFzdC0xZSIsCiAgImRldnBheVByb2R1Y3RDb2RlcyIgOiBudWxsLAogICJtYXJrZXRwbGFjZVByb2R1Y3RDb2RlcyIgOiBudWxsLAogICJ2ZXJzaW9uIiA6ICIyMDE3LTA5LTMwIiwKICAiaW5zdGFuY2VJZCIgOiAiaS0wNTU5NTBiOWNmYmNlMGNlMyIsCiAgImJpbGxpbmdQcm9kdWN0cyIgOiBudWxsLAogICJpbnN0YW5jZVR5cGUiIDogIm00LmxhcmdlIiwKICAicGVuZGluZ1RpbWUiIDogIjIwMTctMTItMjlUMDU6MzQ6MDdaIiwKICAicHJpdmF0ZUlwIiA6ICIxMDAuNjYuMzkuMTU0IiwKICAiYWNjb3VudElkIiA6ICIxNzk3MjcxMDExOTQiLAogICJhcmNoaXRlY3R1cmUiIDogIng4Nl82NCIsCiAgImtlcm5lbElkIiA6IG51bGwsCiAgInJhbWRpc2tJZCIgOiBudWxsLAogICJpbWFnZUlkIiA6ICJhbWktZWExNjU5OTAiLAogICJyZWdpb24iIDogInVzLWVhc3QtMSIKfQAAAAAAADGCAf8wggH7AgEBMGkwXDELMAkGA1UEBhMCVVMxGTAXBgNVBAgTEFdhc2hpbmd0b24gU3RhdGUxEDAOBgNVBAcTB1NlYXR0bGUxIDAeBgNVBAoTF0FtYXpvbiBXZWIgU2VydmljZXMgTExDAgkAsWnMQBVZpBkwDQYJYIZIAWUDBAIBBQCgaTAYBgkqhkiG9w0BCQMxCwYJKoZIhvcNAQcBMBwGCSqGSIb3DQEJBTEPFw0xNzEyMjkwNTM0MTJaMC8GCSqGSIb3DQEJBDEiBCAnRmg+lfohgYY+mmmARsuY+eDhZnxkqSnq3QFOzB0N2jANBgkqhkiG9w0BAQEFAASCAQB8huDo5kOLVGKjGeT3EfK43by0SqS3FThw2VIL4dnFf1OV3c53DT2YjIq4nf9IXFCuE51ch7TLsY9TWTGbLr+U1UF9Y6sahEAfJpp3Vk4KyMwuiELm0/glExksUMgy51tOOAqOMXwoHPBKtnWx9ZVSmv/KUrU3R3armHBrUrcOdZxF2OMPIanWyTJ3aO1uPX4LZOTS/vnM792ED1YlZ1bKl8iu1U0WXEKqc4arkmeqvDhvIs2cg4EKr4QE103dAGCSZzRd+IKd6ZiKivcZP4ul1006IJSrAP0VK4N+PS0bw9xzacMOPPdeP/JiGYlCiiT2tSXwNWB/nKeeSDyPoa4YAAAAAAAA"
const EC2_BAD_DAT = "MIAGCSqGSIb3DQEHAqCAMIACAQExDzANBglghkgBZQMEAgEFADCABgkqhkiG9w0BBwGggCSABIIB1HsKICAiYXZhaWxhYmlsaXR5Wm9uZSIgOiAidXMtZWFzdC0xZSIsCiAgImRldnBheVByb2R1Y3RDb2RlcyIgOiBudWxsLAogICJtYXJrZXRwbGFjZVByb2R1Y3RDb2RlcyIgOiBudWxsLAogICJ2ZXJzaW9uIiA6ICIyMDE3LTA5LTMwIiwKICAiaW5zdGFuY2VJZCIgOiAiaS0wNTU5NTBiOWNmYmNlMGNlMyIsCiAgImJpbGxpbmdQcm9kdWN0cyIgOiBudWxsLAogICJpbnN0YW5jZVR5cGUiIDogIm00LmxhcmdlIiwKICAicGVuZGluZ1RpbWUiIDogIjIwMTctMTItMjlUMDU6MzQ6MDdaIiwKICAicHJpdmF0ZUlwIiA6ICIxMDAuNjYuMzkuMTU0IiwKICAiYWNjb3VudElkIiA6ICIxNzk3MjcxMDExOTQiLAogICJhcmNoaXRlY3R1cmUiIDogIng4Nl82NCIsCiAgImtlcm5lbElkIiA6IG51bGwsCiAgInJhbWRpc2tJZCIgOiBudWxsLAogICJpbWFnZUlkIiA6ICJhbWktMDAwMDAwMDAiLAogICJyZWdpb24iIDogInVzLWVhc3QtMSIKfQAAAAAAADGCAf8wggH7AgEBMGkwXDELMAkGA1UEBhMCVVMxGTAXBgNVBAgTEFdhc2hpbmd0b24gU3RhdGUxEDAOBgNVBAcTB1NlYXR0bGUxIDAeBgNVBAoTF0FtYXpvbiBXZWIgU2VydmljZXMgTExDAgkAsWnMQBVZpBkwDQYJYIZIAWUDBAIBBQCgaTAYBgkqhkiG9w0BCQMxCwYJKoZIhvcNAQcBMBwGCSqGSIb3DQEJBTEPFw0xNzEyMjkwNTM0MTJaMC8GCSqGSIb3DQEJBDEiBCAnRmg+lfohgYY+mmmARsuY+eDhZnxkqSnq3QFOzB0N2jANBgkqhkiG9w0BAQEFAASCAQB8huDo5kOLVGKjGeT3EfK43by0SqS3FThw2VIL4dnFf1OV3c53DT2YjIq4nf9IXFCuE51ch7TLsY9TWTGbLr+U1UF9Y6sahEAfJpp3Vk4KyMwuiELm0/glExksUMgy51tOOAqOMXwoHPBKtnWx9ZVSmv/KUrU3R3armHBrUrcOdZxF2OMPIanWyTJ3aO1uPX4LZOTS/vnM792ED1YlZ1bKl8iu1U0WXEKqc4arkmeqvDhvIs2cg4EKr4QE103dAGCSZzRd+IKd6ZiKivcZP4ul1006IJSrAP0VK4N+PS0bw9xzacMOPPdeP/JiGYlCiiT2tSXwNWB/nKeeSDyPoa4YAAAAAAAA"
const EC2_BAD_SIG = "MIAGCSqGSIb3DQEHAqCAMIACAQExDzANBglghkgBZQMEAgEFADCABgkqhkiG9w0BBwGggCSABIIB1HsKICAiYXZhaWxhYmlsaXR5Wm9uZSIgOiAidXMtZWFzdC0xZSIsCiAgImRldnBheVByb2R1Y3RDb2RlcyIgOiBudWxsLAogICJtYXJrZXRwbGFjZVByb2R1Y3RDb2RlcyIgOiBudWxsLAogICJ2ZXJzaW9uIiA6ICIyMDE3LTA5LTMwIiwKICAiaW5zdGFuY2VJZCIgOiAiaS0wNTU5NTBiOWNmYmNlMGNlMyIsCiAgImJpbGxpbmdQcm9kdWN0cyIgOiBudWxsLAogICJpbnN0YW5jZVR5cGUiIDogIm00LmxhcmdlIiwKICAicGVuZGluZ1RpbWUiIDogIjIwMTctMTItMjlUMDU6MzQ6MDdaIiwKICAicHJpdmF0ZUlwIiA6ICIxMDAuNjYuMzkuMTU0IiwKICAiYWNjb3VudElkIiA6ICIxNzk3MjcxMDExOTQiLAogICJhcmNoaXRlY3R1cmUiIDogIng4Nl82NCIsCiAgImtlcm5lbElkIiA6IG51bGwsCiAgInJhbWRpc2tJZCIgOiBudWxsLAogICJpbWFnZUlkIiA6ICJhbWktZWExNjU5OTAiLAogICJyZWdpb24iIDogInVzLWVhc3QtMSIKfQAAAAAAADGCAf8wggH7AgEBMGkwXDELMAkGA1UEBhMCVVMxGTAXBgNVBAgTEFdhc2hpbmd0b24gU3RhdGUxEDAOBgNVBAcTB1NlYXR0bGUxIDAeBgNVBAoTF0FtYXpvbiBXZWIgU2VydmljZXMgTExDAgkAsWnMQBVZpBkwDQYJYIZIAWUDBAIBBQCgaTAYBgkqhkiG9w0BCQMxCwYJKoZIhvcNAQcBMBwGCSqGSIb3DQEJBTEPFw0xNzEyMjkwNTM0MTJaMC8GCSqGSIb3DQEJBDEiBCAnRmg+lfohgYY+mmmARsuY+eDhZnxkqSnq3QFOzB0N2jANBgkqhkiG9w0BAQEFAASCAQB9huDo5kOLVGKjGeT3EfK43by0SqS3FThw2VIL4dnFf1OV3c53DT2YjIq4nf9IXFCuE51ch7TLsY9TWTGbLr+U1UF9Y6sahEAfJpp3Vk4KyMwuiELm0/glExksUMgy51tOOAqOMXwoHPBKtnWx9ZVSmv/KUrU3R3armHBrUrcOdZxF2OMPIanWyTJ3aO1uPX4LZOTS/vnM792ED1YlZ1bKl8iu1U0WXEKqc4arkmeqvDhvIs2cg4EKr4QE103dAGCSZzRd+IKd6ZiKivcZP4ul1006IJSrAP0VK4N+PS0bw9xzacMOPPdeP/JiGYlCiiT2tSXwNWB/nKeeSDyPoa4YAAAAAAAA"

func TestVerifyAttestation(t *testing.T) {

	provider := &AwsCloudProvider{
		myAmi: EC2_METADATA_AMI,
	}

	cmsBytes, err := base64.StdEncoding.DecodeString(EC2_RSA2048)
	if err != nil {
		t.Fatal(err)
	}

	err = provider.VerifyAttestation(cmsBytes)
	if err != nil {
		t.Fatal("Failed to verify a valid attestation")
	}

	provider.myAmi = "ami-00000000"
	err = provider.VerifyAttestation(cmsBytes)
	if err == nil || err.Error() != "Client image id ami-ea165990 doesn't match instance image id ami-00000000" {
		t.Fatalf("Got invalid error with mismatched AMI: %v", err)
	}

	cmsBytes, err = base64.StdEncoding.DecodeString(EC2_BAD_DAT)
	if err != nil {
		t.Fatal(err)
	}
	err = provider.VerifyAttestation(cmsBytes)
	if err == nil || !strings.HasPrefix(err.Error(), "pkcs7: Message digest mismatch") {
		t.Fatalf("Got invalid error with bad data on CMS: %v", err)
	}

	cmsBytes, err = base64.StdEncoding.DecodeString(EC2_BAD_SIG)
	if err != nil {
		t.Fatal(err)
	}
	err = provider.VerifyAttestation(cmsBytes)
	if err == nil || !strings.HasPrefix(err.Error(), "crypto/rsa: verification error") {
		t.Fatalf("Got invalid error with bad signature on CMS: %v", err)
	}
}

package lib

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"log"
	"net/http"
	"strings"
)

func QueryCtrl(endpoint string, dataBin string) (node *xmlquery.Node, err error) {
	req, _ := http.NewRequest(
		"POST",
		endpoint+"/YamahaRemoteControl/ctrl",
		strings.NewReader(dataBin),
	)
	log.Println(req.Method, req.URL, dataBin)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Error during the request: " + err.Error() + ". Is the endpoint correct?")
	}
	doc, err := xmlquery.Parse(resp.Body)
	if err != nil {
		return nil, errors.New("Error parsing the response: " + err.Error() + ". Is the endpoint correct?")
	}
	return doc, nil
}

func GetParams(endpoint string) (node *xmlquery.Node, err error) {
	return QueryCtrl(
		endpoint,
		`<YAMAHA_AV cmd="GET"><Main_Zone><Basic_Status>GetParam</Basic_Status></Main_Zone></YAMAHA_AV>`)
}

func SetSource(endpoint string, newSource string) (node *xmlquery.Node, err error) {
	buf := bytes.NewBufferString("")
	_ = xml.EscapeText(buf, []byte(newSource))
	escaped := buf.String()
	return QueryCtrl(
		endpoint,
		fmt.Sprintf(
			`<YAMAHA_AV cmd="PUT"><Main_Zone><Input><Input_Sel>%s</Input_Sel></Input></Main_Zone></YAMAHA_AV>`,
			escaped))
}

func SetPower(endpoint string, power string) (node *xmlquery.Node, err error) {
	buf := bytes.NewBufferString("")
	_ = xml.EscapeText(buf, []byte(power))
	escaped := buf.String()
	return QueryCtrl(
		endpoint,
		fmt.Sprintf(
			`<YAMAHA_AV cmd="PUT"><Main_Zone><Power_Control><Power>%s</Power></Power_Control></Main_Zone></YAMAHA_AV>`,
			escaped))
}

func PowerOn(endpoint string) (node *xmlquery.Node, err error) {
	return SetPower(endpoint, "On")
}

func PowerOff(endpoint string) (node *xmlquery.Node, err error) {
	return SetPower(endpoint, "Standby")
}

func SetVolume(endpoint string, volume int) (node *xmlquery.Node, err error) {
	return QueryCtrl(
		endpoint,
		fmt.Sprintf(
			`<YAMAHA_AV cmd="PUT"><Main_Zone><Volume><Lvl><Val>%d</Val><Exp>1</Exp><Unit>dB</Unit></Lvl></Volume></Main_Zone></YAMAHA_AV>`,
			volume))
}

func SetMute(endpoint string, muteValue string) (node *xmlquery.Node, err error) {
	buf := bytes.NewBufferString("")
	_ = xml.EscapeText(buf, []byte(muteValue))
	escaped := buf.String()
	return QueryCtrl(
		endpoint,
		fmt.Sprintf(
			`<YAMAHA_AV cmd="PUT"><Main_Zone><Volume><Mute>%s</Mute></Volume></Main_Zone></YAMAHA_AV>`,
			escaped))
}

func Mute(endpoint string) (node *xmlquery.Node, err error) {
	return SetMute(endpoint, "On")
}

func Unmute(endpoint string) (node *xmlquery.Node, err error) {
	return SetMute(endpoint, "Off")
}

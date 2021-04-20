package mansekum

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Credential ...
type Credential struct {
	BaseURL  string
	Username string
	Password string
	AppCode  string
	AdminID  string
	AdminKey string
	AppID    string
}

const (
	dukcapilURL       = "/dukcapil/api/nik"
	insertUserURL     = "/UM/insertuser.prc"
	checkEmailURL     = "/UM/checkemail.prc"
	loginURL          = "/UM/login.prc"
	reqResetPassURL   = "/UM/resetpwd.prc"
	sendSmsURL        = "/sms/api/sms/send/"
	editUserURL       = "/UM/edituser.prc"
	openingStatusURL  = "/OA/openingstatus.inq"
	userInfoURL       = "/UM/userinfo.inq"
	resetPassURL      = "/UM/setpwd.prc"
	resetPinURL       = "/UM/setpin.prc"
	getUserByPhoneURL = "/UM/usersbyphone.inq"
	cashInfoURL       = "/INFO/cash.inq"
)

// GetDukcapilData ...
func (cred *Credential) GetDukcapilData(identityNumber string) (res Response, err error) {
	// 3502036305910001
	// 3275056305940005
	// 3173073007790007
	// 3603125510910006
	// 3302251706880001
	// 3373031702990001

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + dukcapilURL + "?y=" + identityNumber

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("GET", fullURL, nil)
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)

	return res, err
}

// InsertUser ...
func (cred *Credential) InsertUser(name, email, phone, password string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + insertUserURL

	// Make payload
	payload := map[string]interface{}{
		"application": cred.AppCode,
		"userAdmin":   cred.AdminID,
		"key":         cred.AdminKey,
		"name":        name,
		"alias":       name,
		"email":       email,
		"phone":       phone,
		"pwd":         password,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// CheckEmail ...
func (cred *Credential) CheckEmail(email string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + checkEmailURL

	// Make payload
	payload := map[string]interface{}{
		"email": email,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// Login ...
func (cred *Credential) Login(email, password string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + loginURL

	// Make payload
	payload := map[string]interface{}{
		"application": cred.AppCode,
		"id":          email,
		"pwd":         password,
		"ipaddress":   "127.0.0.1",
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// ReqResetPassword ...
func (cred *Credential) ReqResetPassword(id string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + reqResetPassURL

	// Make payload
	payload := map[string]interface{}{
		"application": cred.AppCode,
		"userAdmin":   cred.AdminID,
		"key":         cred.AdminKey,
		"id":          id,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// SendSms ...
func (cred *Credential) SendSms(phone, message, channel string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + sendSmsURL + channel

	// Make payload
	payload := map[string]interface{}{
		"destination": "+" + strings.Replace(phone, "-", "", -1),
		"text":        message,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// EditUser ...
func (cred *Credential) EditUser(id, name, email, phone string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + editUserURL

	// Make payload
	payload := map[string]interface{}{
		"application": cred.AppCode,
		"userAdmin":   cred.AdminID,
		"key":         cred.AdminKey,
		"id":          id,
		"name":        name,
		"alias":       name,
		"email":       email,
		"phone":       phone,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// OpeningStatus ...
func (cred *Credential) OpeningStatus(userID string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + openingStatusURL

	// Make payload
	payload := map[string]interface{}{
		"ID": userID,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// UserInfo ...
func (cred *Credential) UserInfo(userID string) (res UserInfo, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + userInfoURL

	// Make payload
	payload := map[string]interface{}{
		"id": userID,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// ResetPassword ...
func (cred *Credential) ResetPassword(userID, password string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + resetPassURL

	// Make payload
	payload := map[string]interface{}{
		"id":  userID,
		"pwd": password,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// ResetPin ...
func (cred *Credential) ResetPin(userID, pin string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + resetPinURL

	// Make payload
	payload := map[string]interface{}{
		"id":  userID,
		"pin": pin,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// GetUserByPhone ...
func (cred *Credential) GetUserByPhone(phone string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + getUserByPhoneURL

	// Make payload
	payload := map[string]interface{}{
		"phone": phone,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	// Remove \n in response string
	bodyString := strings.Replace(string(body), "\n", "", 1)
	body = []byte(bodyString)
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

// CashInfo ...
func (cred *Credential) CashInfo(userID string) (res map[string]interface{}, err error) {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(cred.Username+":"+cred.Password))
	fullURL := cred.BaseURL + cashInfoURL

	// Make payload
	payload := map[string]interface{}{
		"id": userID,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return res, errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(pBody))
	r.Header.Add("Authorization", auth)
	r.Header.Add("X-AppId", cred.AppID)
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(fullURL + " " + string(body))
	}

	return res, err
}

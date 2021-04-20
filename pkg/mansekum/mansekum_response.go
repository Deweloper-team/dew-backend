package mansekum

// Response ...
type Response struct {
	Content []Contect `json:"content"`
}

// Laki-Laki Perempuan
// Contect ...
type Contect struct {
	Agama       string `json:"AGAMA"`
	Alamat      string `json:"ALAMAT"`
	Dusun       string `json:"DUSUN"`
	EktpCreated string `json:"EKTP_CREATED"`
	EktpStatus  string `json:"EKTP_STATUS"`
	GolDarah    string `json:"GOL_DARAH"`
	JenisKlmin  string `json:"JENIS_KLMIN"`
	JenisPkrjn  string `json:"JENIS_PKRJN"`
	KabName     string `json:"KAB_NAME"`
	KecName     string `json:"KEC_NAME"`
	KelName     string `json:"KEL_NAME"`
	KodePos     int    `json:"KODE_POS"`
	LastUpdated string `json:"LAST_UPDATED"`
	NamaLgkp    string `json:"NAMA_LGKP"`
	NamaLgkpIbu string `json:"NAMA_LGKP_IBU"`
	Nik         int    `json:"NIK"`
	NoKab       int    `json:"NO_KAB"`
	NoKec       int    `json:"NO_KEC"`
	NoKel       int    `json:"NO_KEL"`
	NoKk        int    `json:"NO_KK"`
	NoProp      int    `json:"NO_PROP"`
	NoRt        int    `json:"NO_RT"`
	NoRw        int    `json:"NO_RW"`
	PddkAkh     string `json:"PDDK_AKH"`
	PropName    string `json:"PROP_NAME"`
	Statuskawin string `json:"STATUS_KAWIN"`
	StatHbkel   string `json:"STAT_HBKEL"`
	TglLhr      string `json:"TGL_LHR"`
	TmptLhr     string `json:"TMPT_LHR"`
}

// UserInfo ...
type UserInfo struct {
	Code    string           `json:"code"`
	Objects []UserInfoObject `json:"objects"`
	Message string           `json:"message"`
	Status  string           `json:"status"`
}

// UserInfoObject ...
type UserInfoObject struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	MotherName string `json:"mother_name"`
	Birthdate  string `json:"birthdate"`
}

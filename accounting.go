package accounting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/leekchan/accounting"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/frontpage", frontPage)
	http.HandleFunc("/pengeluaran", pagePengeluaran)
	http.HandleFunc("/pengeluaran-refresh", pagePengeluaranRefresh)
	http.HandleFunc("/pemasukan", pagePemasukan)
	http.HandleFunc("/hutang", pageHutang)
	http.HandleFunc("/tukar-jaga", pageTukarJaga)
	http.HandleFunc("/pengaturan", pagePengaturan)
	http.HandleFunc("/add-pengeluaran", addPengeluaran)
	http.HandleFunc("/add-pemasukan", addPemasukan)
	http.HandleFunc("/createkursor", createKursor)
	http.HandleFunc("/getbulan", getBulan)
	// http.HandleFunc("/balance-resume", balanceResume)
	http.HandleFunc("/create-resume", balanceResume)
	http.HandleFunc("/delete-pengeluaran", deletePengeluaran)
	http.HandleFunc("/get-pemasukan-bulan", getPemasukanBulan)
	http.HandleFunc("/delete-pemasukan", deletePemasukan)
	// http.HandleFunc("/tambah-piutang", tambahPiutang)
	// http.HandleFunc("/get-piutang", getPiutang)
	// http.HandleFunc("/tambah-hutang", tambahHutang)
	http.HandleFunc("/tambah-tukar-jaga", TambahTukarJaga)
	http.HandleFunc("/get-tukar-jaga", getTukarJaga)
	http.HandleFunc("/ubah-tanggal-bayar-jaga", ubahTglBayarJaga)
}

// func tambahHutang(w http.ResponseWriter, r *http.Request) {
// 	ctx := appengine.NewContext(r)
// 	dat := &CatchDataJson{}
// 	err := json.NewDecoder(r.Body).Decode(dat)
// 	if err != nil {
// 		ErrorRec(ctx, "Gagal mengambil dat dari klien", err)
// 		return
// 	}
// 	defer r.Body.Close()
// 	// log.Infof(ctx, "isi data adaah: %v", dat)

// }
// func getPiutang(w http.ResponseWriter, r *http.Request) {
// 	ctx := appengine.NewContext(r)
// 	q := datastore.NewQuery("Piutang").Filter("Status >", 0)
// 	list := []Piutang{}
// 	_, err := q.GetAll(ctx, &list)
// 	if err != nil {
// 		ErrorRec(ctx, "gagal mengambil data", err)
// 		SendBackError(w, "gagal mengambil data", 500)
// 		return
// 	}

// 	// for {
// 	// 	pi := &Piutang{}
// 	// 	_, err := t.Next(pi)
// 	// 	if err == datastore.Done {
// 	// 		break
// 	// 	}
// 	// 	list = append(list, *pi)
// 	// }
// 	// log.Infof(ctx, "List adaalh: %v", list)
// 	w.WriteHeader(200)
// 	fmt.Fprint(w, GenTemplate(ctx, list, "hal-piutang-tabel"))
// }
// func tambahPiutang(w http.ResponseWriter, r *http.Request) {
// 	ctx := appengine.NewContext(r)
// 	piu := &CatchPiutang{}
// 	json.NewDecoder(r.Body).Decode(piu)
// 	defer r.Body.Close()
// 	log.Infof(ctx, "jumlah adalah: %v", piu.Jumlah)
// 	k := datastore.NewIncompleteKey(ctx, "Piutang", nil)
// 	inp := &Piutang{
// 		TanggalInput: timeNowIndonesia(),
// 		NamaJenis:    piu.Nama,
// 		// "0" = Lunas, "1" = Belum lunas, "2" = Lunas sebagian
// 		Status: 1,
// 		Link:   k.Encode(),
// 	}
// 	if piu.Tanggal != "" {
// 		inp.TanggalPiutang = ChangeStringtoTime(piu.Tanggal)
// 		// log.Infof(ctx, "Tanggal adalah: %v", inp.TanggalPiutang)
// 	}
// 	if piu.Tanggal == "" {
// 		jml, _ := strconv.Atoi(piu.Jumlah)
// 		inp.Jumlah = jml
// 	}

// 	_, err := datastore.Put(ctx, k, inp)
// 	if err != nil {
// 		ErrorRec(ctx, "Gagal menyimpan piutang", err)
// 		SendBackError(w, "Gagal Menyimpan Piutang", 500)
// 	}
// 	SendBackSuccess(w, nil, "", "", "")
// }
func deletePemasukan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	in := &Input{}
	json.NewDecoder(r.Body).Decode(in)
	r.Body.Close()
	k, _ := datastore.DecodeKey(in.Link)
	err := datastore.Delete(ctx, k)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Kesalahan di server, gagal menghapus data")
		return
	}
	SendBackSuccess(w, nil, "Berhasil menghapus data", "", "")
}
func getPemasukanBulan(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	in := &Input{}
	json.NewDecoder(r.Body).Decode(in)
	defer r.Body.Close()
	kur := &Kursor{}
	err := json.Unmarshal([]byte(in.NamaItem), kur)
	if err != nil {
		ErrorRec(ctx, "Gagal membaca json", err)
		return
	}
	// list := getMonthly(ctx, kur.Link, kur.Point)
	pg := resumePemasukan(ctx, getMonthly(ctx, kur.Link, kur.Point))
	res := &ResponseJson{
		ScriptTambahan: RupiahString(pg.PemasukanBulanIni),
		Script:         GenTemplate(ctx, pg.DaftarPemasukan, "hal-pemasukan-content"),
	}
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
	// json.NewEncoder()

}
func deletePengeluaran(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	// link := r.FormValue("link")
	in := &Input{}
	json.NewDecoder(r.Body).Decode(in)
	defer r.Body.Close()
	// log.Infof(ctx, "Link adalah: %v", in.Link)
	key, _ := datastore.DecodeKey(in.Link)
	// log.Infof(ctx, "Key adalah: %v", key)
	err := datastore.Delete(ctx, key)
	if err != nil {
		ErrorRec(ctx, "Gagal menghapus entry", err)
		w.WriteHeader(422)
		fmt.Fprintf(w, "Gagal menghapus entri")
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Berhasil menghapus entri")
	}
}

func createBalanceResume(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	kur := getKursor(ctx)
	for _, v := range kur {
		list := getMonthly(ctx, v.Link, v.Point)
		pem := 0
		i := &pem
		peng := 0
		h := &peng
		for _, n := range list {
			if n.Pemasukan == "" {
				exp, _ := strconv.Atoi(n.Pengeluaran)
				*h = *h + exp
			}
			if n.Pengeluaran == "" {
				inc, _ := strconv.Atoi(n.Pemasukan)
				*i = *i + inc
			}
		}
		bal := &ResumePerbulan{
			Tanggal:     timeNowIndonesia(),
			Pengeluaran: peng,
			Pemasukan:   pem,
		}
		_, err := datastore.Put(ctx, datastore.NewKey(ctx, "BalanceResume", v.Point, 0, nil), bal)
		if err != nil {
			ErrorRec(ctx, "Gagal Menyimpan Resume Balance", err)
			return
		}
	}
}
func balanceResume(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	sk := timeNowIndonesia().AddDate(0, -1, 0)
	peng := 0
	i := &peng
	pem := 0
	h := &pem
	q := datastore.NewQuery("Balance").Order("-Tanggal")
	t := q.Run(ctx)
	for {
		j := &Input{}
		_, err := t.Next(j)
		if err == datastore.Done {
			log.Infof(ctx, "End of data")
		}
		if j.Tanggal.In(ZonaIndo()).Before(sk) {
			break
		}
		if j.Pengeluaran == "" {
			deb, _ := strconv.Atoi(j.Pemasukan)
			*h = *h + deb
		}
		if j.Pemasukan == "" {
			kred, _ := strconv.Atoi(j.Pengeluaran)
			*i = *i + kred
		}
	}
	ume := getResume(ctx)
	res := &ResumePerbulan{
		Tanggal:          timeNowIndonesia(),
		Pemasukan:        *h,
		Pengeluaran:      *i,
		TotalPemasukan:   ume.TotalPemasukan + *h,
		TotalPengeluaran: ume.TotalPengeluaran + *i,
	}
	_, err := datastore.Put(ctx, datastore.NewKey(ctx, "ResumePerbulan", sk.Format("2006/01"), 0, nil), res)
	if err != nil {
		ErrorRec(ctx, "Gagal Menyimpan Resume Perbulan", err)
		return
	}
}

func getResume(c context.Context) ResumePerbulan {
	q := datastore.NewQuery("ResumePerbulan").Order("-Tanggal")
	t := q.Run(c)
	j := &ResumePerbulan{}
	for {
		_, err := t.Next(j)
		if err == datastore.Done {
			break
		}
		break
	}
	return *j
}

func getBulan(w http.ResponseWriter, r *http.Request) {
	// kur := r.FormValue("link")
	// tgl := r.FormValue("tgl")
	in := &Input{}
	json.NewDecoder(r.Body).Decode(in)
	defer r.Body.Close()
	ctx := appengine.NewContext(r)
	list := getMonthly(ctx, in.Link, in.NamaItem)
	// log.Infof(ctx, "List adalah: %v", list)
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	js := &ResponseJson{
		Script: GenTemplate(ctx, resumePengeluaran(ctx, list), "hal-pengeluaran-content"),
	}
	json.NewEncoder(w).Encode(js)
}
func createKursor(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	q := datastore.NewQuery("Balance").Order("-Tanggal")
	t := q.Run(ctx)
	in := &Input{}
	wkt := timeNowIndonesia()
	zone, _ := time.LoadLocation("Asia/Makassar")
	bln := time.Date(wkt.Year(), wkt.Month(), 1, 0, 0, 0, 0, zone)
	tgl := bln.AddDate(0, -1, 0).Format("2006/01")
	keyKur := datastore.NewKey(ctx, "KursorBalance", tgl, 0, nil)
	for {
		t.Next(in)
		if in.Tanggal.In(zone).Before(bln) == true {
			cursor, _ := t.Cursor()
			kur := &Kursor{
				Point: cursor.String(),
			}
			if _, err := datastore.Put(ctx, keyKur, kur); err != nil {
				ErrorRec(ctx, "Gagal membuat kursor", err)
			}
			break
		}
	}
}
func ZonaIndo() *time.Location {
	zone, _ := time.LoadLocation("Asia/Makassar")
	return zone
}
func CountExpense(c context.Context, kur, tgl string) int {
	if kur == "" {
		skrng := time.Date(timeNowIndonesia().Year(), timeNowIndonesia().Month(), 1, 0, 0, 0, 0, ZonaIndo())
		q := datastore.NewQuery("Balance").Order("-Tanggal")
		t := q.Run(c)
		in := &Input{}
		peng := 0
		i := &peng
		for {
			_, err := t.Next(in)
			if err == datastore.Done {
				// log.Infof(ctx, "Data habis")
				break
			}
			if err != nil {
				ErrorRec(c, "Gagal mengambil data", err)
				// js.ModalScript = "Gagal mengambil data"
				// json.NewEncoder(w).Encode(js)
				break
			}
			if in.Tanggal.In(ZonaIndo()).Before(skrng) == true {
				// log.Infof(ctx, "Tanggal Kadaluarsa")
				break
			}
			// log.Infof(ctx, "Tanggal data adalH: %v", in.Tanggal)
			kel, _ := strconv.Atoi(in.Pengeluaran)
			*i = *i + kel
		}
		return peng
	} else {
		list := getMonthly(c, kur, tgl)
		peng := 0
		i := &peng
		for _, v := range list {
			kel, _ := strconv.Atoi(v.Pengeluaran)
			*i = *i + kel
		}
		return peng
	}
}

func CountBalance(c context.Context) int {
	maks := &MaksPengeluaran{}
	p := datastore.NewQuery("MaksPengeluaran").Order("-Tanggal")
	s := p.Run(c)
	for {
		s.Next(maks)
		break
	}
	expense := CountExpense(c, "", "")
	return maks.MaksPengeluaran - expense
}
func frontPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	maks := &MaksPengeluaran{}
	p := datastore.NewQuery("MaksPengeluaran")
	s := p.Run(ctx)
	for {
		s.Next(maks)
		break
	}
	skrng := time.Date(timeNowIndonesia().Year(), timeNowIndonesia().Month(), 1, 0, 0, 0, 0, ZonaIndo())
	// log.Infof(ctx, "Skarang adalh: %v", skrng)
	q := datastore.NewQuery("Balance").Order("-Tanggal")
	t := q.Run(ctx)
	in := &Input{}
	cptl := maks.MaksPengeluaran
	js := &ResponseJson{}
	i := &cptl
	for {
		_, err := t.Next(in)
		if err == datastore.Done {
			// log.Infof(ctx, "Data habis")
			break
		}
		if err != nil {
			ErrorRec(ctx, "Gagal mengambil data", err)
			js.ModalScript = "Gagal mengambil data"
			json.NewEncoder(w).Encode(js)
			break
		}
		if in.Tanggal.In(ZonaIndo()).Before(skrng) == true {
			// log.Infof(ctx, "Tanggal Kadaluarsa")
			break
		}
		// log.Infof(ctx, "Tanggal data adalH: %v", in.Tanggal)
		peng, _ := strconv.Atoi(in.Pengeluaran)
		*i = *i - peng
	}
	js.Data = cptl
	maks.MaksPengeluaran = cptl
	js.Script = GenTemplate(ctx, maks, "main-content")
	json.NewEncoder(w).Encode(js)
}
func addPemasukan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	inp := &Input{
		Tanggal: timeNowIndonesia(),
	}
	json.NewDecoder(r.Body).Decode(inp)
	if inp.NamaItem == "" || inp.Pemasukan == "" {
		ErrorRec(ctx, "Kolom isian kosong", nil)
		w.WriteHeader(406)
		fmt.Fprint(w, "Kolom isian kosong!")
		return
	}
	defer r.Body.Close()
	inputKey := datastore.NewIncompleteKey(ctx, "Balance", nil)
	_, err := datastore.Put(ctx, inputKey, inp)
	if err != nil {
		ErrorRec(ctx, "menyimpan data pemasukan", err)
		w.WriteHeader(422)
		fmt.Fprint(w, "Gagal Menyimpan Data. Input data kembali")
		return
	}
	inp.ServerMsg = "Berhasil Menambahkan Item"
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(inp)
}
func addPengeluaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	inp := &Input{
		Tanggal: timeNowIndonesia(),
	}
	json.NewDecoder(r.Body).Decode(inp)
	if inp.NamaItem == "" || inp.Pengeluaran == "" {
		ErrorRec(ctx, "Kolom isian kosong", nil)
		w.WriteHeader(406)
		fmt.Fprint(w, "Kolom isian kosong!")
		return
	}
	defer r.Body.Close()
	inputKey := datastore.NewIncompleteKey(ctx, "Balance", nil)
	_, err := datastore.Put(ctx, inputKey, inp)
	if err != nil {
		ErrorRec(ctx, "menyimpan data pengeluaran", err)
		w.WriteHeader(422)
		fmt.Fprint(w, "Gagal Menyimpan Data. Input data kembali")
		return
	}
	inp.ServerMsg = "Berhasil Menambahkan Item"
	rp := CountBalance(ctx)
	// rp, _ := strconv.Atoi(inp.Pengeluaran)
	inp.Pemasukan = RupiahString(rp)
	inp.Pengeluaran = strconv.Itoa(rp)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(inp)
	// fmt.Fprint(w, "Berhasil menambahkan")
}
func timeNowIndonesia() time.Time {
	zone, _ := time.LoadLocation("Asia/Makassar")
	now := time.Now()
	return now.In(zone)
}
func getKursor(c context.Context) []Kursor {
	ku := datastore.NewQuery("KursorBalance")
	kur := []Kursor{}
	kurs := []Kursor{}
	keyKurs, err := ku.GetAll(c, &kur)
	if err != nil || err == datastore.ErrNoSuchEntity {
		log.Errorf(c, "Tidak dapat mengambil kursor")
		return nil
	} else {
		for _, m := range kur {
			for _, g := range keyKurs {

				b := Kursor{
					// Point: "<a href='' class='w3-bar-item w3-button' data-link='" + m.Point + "'>" + g.StringID() + "</a>",
					Point: g.StringID(),
					Link:  m.Point,
				}
				kurs = append(kurs, b)
			}
		}
	}
	return kurs
}

func getListItem(c context.Context, q *datastore.Query, tgl time.Time) []Input {
	peng := []Input{}
	t := q.Run(c)
	for {
		j := &Input{}
		k, err := t.Next(j)
		if err == datastore.Done {
			break
		}
		if j.Tanggal.In(ZonaIndo()).Before(tgl) == true {
			break
		}
		j.Link = k.Encode()
		peng = append(peng, *j)
	}
	return peng
}
func getMonthly(c context.Context, kur, tgl string) []Input {
	s := timeNowIndonesia()
	if kur == "" {
		q := datastore.NewQuery("Balance").Order("-Tanggal")
		return getListItem(c, q, time.Date(s.Year(), s.Month(), 1, 0, 0, 0, 0, ZonaIndo()))
	} else {
		q := datastore.NewQuery("Balance").Order("-Tanggal")
		kurs, _ := datastore.DecodeCursor(kur)
		q = q.Start(kurs)
		yr, _ := time.Parse("2006/01/02", tgl+"/01")
		// log.Infof(c, "Bulan adalah: %v", yr)
		return getListItem(c, q, yr.In(ZonaIndo()))
	}
}
func resumePengeluaran(c context.Context, peng []Input) PengeluaranPage {
	total := 0
	ftot := &total
	list := []Input{}
	for _, v := range peng {
		if v.Pengeluaran == "" {
			continue
		}
		kel, _ := strconv.Atoi(v.Pengeluaran)
		*ftot = *ftot + kel
		list = append(list, v)
	}

	pg := &PengeluaranPage{
		TotalPengeluaran:  total,
		DaftarPengeluaran: list,
		DaftarKursor:      getKursor(c),
	}
	return *pg
}
func resumePemasukan(c context.Context, pem []Input) PemasukanPage {
	pema := 0
	i := &pema
	j := []Input{}
	for _, v := range pem {
		if v.Pemasukan == "" {
			continue
		}
		j = append(j, v)
		inc, _ := strconv.Atoi(v.Pemasukan)
		*i = *i + inc
	}
	pg := &PemasukanPage{
		DaftarPemasukan:   j,
		PemasukanBulanIni: pema,
		DaftarKursor:      getKursor(c),
	}
	return *pg
}
func pagePengeluaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	peng := getMonthly(ctx, "", "")
	pg := resumePengeluaran(ctx, peng)
	w.WriteHeader(200)
	fmt.Fprint(w, GenTemplate(ctx, pg, "hal-pengeluaran", "hal-pengeluaran-navbar", "hal-pengeluaran-content"))
}
func pagePengeluaranRefresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	peng := getMonthly(ctx, "", "")
	pg := resumePengeluaran(ctx, peng)
	// log.Infof(ctx, "html adalah: %v", GenTemplate(ctx, pg, "hal-pengeluaran-content"))
	w.WriteHeader(200)
	fmt.Fprint(w, GenTemplate(ctx, pg, "hal-pengeluaran-content"))
}

func pagePemasukan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	list := getMonthly(ctx, "", "")
	res := resumePemasukan(ctx, list)
	sume := getResume(ctx)
	res.DaftarKursor = getKursor(ctx)
	res.TotalTabungan = sume.TotalPemasukan + res.PemasukanBulanIni - sume.TotalPengeluaran
	page := GenTemplate(ctx, res, "hal-pemasukan")
	fmt.Fprint(w, page)
}

func pageHutang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	page := GenTemplate(ctx, nil, "hal-hutang")
	fmt.Fprint(w, page)
}

func pageTukarJaga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	page := GenTemplate(ctx, nil, "hal-tukar-jaga")
	fmt.Fprint(w, page)
}
func pagePengaturan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	page := GenTemplate(ctx, nil, "hal-pengaturan")
	fmt.Fprint(w, page)
}
func responsePost(w http.ResponseWriter, res *ResponseJson) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, "<a href='%s'>Sign in or Register</a>", url)
		return
	}
	logout, _ := user.LogoutURL(ctx, "/")
	if u.Email != "suryasedana@gmail.com" {
		http.Redirect(w, r, logout, 403)
	} else {
		me := FrontPage{
			LogOut:   logout,
			UserName: "I Wayan Surya Sedana",
		}
		front := GenTemplate(ctx, me, "index")
		fmt.Fprint(w, front)
	}
}

func GenTemplate(c context.Context, n interface{}, temp ...string) string {
	b := new(bytes.Buffer)
	funcs := template.FuncMap{
		"jam": func(t time.Time) string {
			zone, _ := time.LoadLocation("Asia/Makassar")
			return t.In(zone).Format("15:04")
		},
		"inc": func(i int) int {
			return i + 1
		},
		"rp": func(i int) string {
			ac := accounting.Accounting{
				Symbol:    "Rp ",
				Precision: 2,
				Thousand:  ".",
				Decimal:   ",",
			}
			m := fmt.Sprint(ac.FormatMoney(i))
			return m
		},
		"rpi": func(s string) string {
			rp, _ := strconv.Atoi(s)
			ac := accounting.Accounting{
				Symbol:    "Rp ",
				Precision: 2,
				Thousand:  ".",
				Decimal:   ",",
			}
			m := fmt.Sprint(ac.FormatMoney(rp))
			return m
		},
		"kur": func(kur Kursor) string {
			js, _ := json.Marshal(kur)
			return string(js)
		},
		"strtgl": func(t time.Time) string {
			return t.Format("Mon, 02/01/2006")
		},
		"istimezero": func(t time.Time) bool {
			return t.IsZero()
		},
		"convstrjaga": func(j string) string {
			var m string
			switch j {
			case "1":
				m = "Pagi"
			case "2":
				m = "Sore"
			case "3":
				m = "Malam"
			}
			return m
		},
	}

	tmpl := template.New("")
	for k, v := range temp {
		if k == 0 {
			tmp := template.Must(template.New(v + ".html").Funcs(funcs).ParseFiles("templates/" + v + ".html"))
			tmpl = tmp
		}
	}

	for k, v := range temp {
		if k != 0 {
			temp, err := template.Must(tmpl.Clone()).ParseFiles("templates/" + v + ".html")
			if err != nil {
				ErrorRec(c, "parse template multiple", err)
				return ""
			}
			tmpl = temp
		}
	}
	err := tmpl.Execute(b, n)
	if err != nil {
		ErrorRec(c, "eksekusi template", err)
		return ""
	}

	return b.String()
}

func ErrorRec(c context.Context, topik string, err error) {
	msg := "Telah terjadi kesalahan dalam " + topik + " : %v"
	log.Errorf(c, msg, err)
}

func RupiahString(i int) string {
	ac := accounting.Accounting{
		Symbol:    "Rp ",
		Precision: 2,
		Thousand:  ".",
		Decimal:   ",",
	}
	m := fmt.Sprint(ac.FormatMoney(i))
	return m
}

func ChangeStringtoTime(tgl string) time.Time {
	str, _ := time.ParseInLocation("2006-1-02", tgl, ZonaIndo())
	return str
}

func SendBackError(w http.ResponseWriter, t string, n int) {
	w.WriteHeader(n)
	fmt.Fprint(w, t)
}

func SendBackSuccess(w http.ResponseWriter, dat interface{}, script, modal, tambahan string) {
	w.WriteHeader(200)
	res := &ResponseJson{
		Data:           dat,
		Script:         script,
		ModalScript:    modal,
		ScriptTambahan: tambahan,
	}
	json.NewEncoder(w).Encode(res)
}

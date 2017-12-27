package accounting

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	_ "html/template"
	"net/http"
	_ "strconv"
	_ "time"

	_ "github.com/leekchan/accounting"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	_ "google.golang.org/appengine/user"
)

func TambahTukarJaga(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	dat := &CatchDataJson{}
	json.NewDecoder(r.Body).Decode(dat)
	// log.Infof(ctx, "Isi data adalah: %v", dat)
	k := datastore.NewIncompleteKey(ctx, "TukarJaga", nil)
	jag := &TukarJaga{
		TanggalInput:      timeNowIndonesia(),
		Nama:              dat.Data1,
		JagaHutang:        dat.Data2,
		TanggalJagaHutang: ChangeStringtoTime(dat.Data3),
		JagaBayar:         dat.Data4,
		TanggalJagaBayar:  ChangeStringtoTime(dat.Data5),
		Link:              k.Encode(),
		Status:            "1",
	}
	_, err := datastore.Put(ctx, k, jag)
	if err != nil {
		ErrorRec(ctx, "Gagal Menambahkan Data", err)
		SendBackError(w, "Gagal Menambahkan data", 500)
		return
	}
	list, err := getListTukarJaga(ctx)
	if err != nil {
		ErrorRec(ctx, "Gagal mengambil list", err)
		SendBackError(w, "Gagal mengambil list", 500)
		return
	}
	list = append([]TukarJaga{*jag}, list...)
	// q := datastore.NewQuery("TukarJaga").Order("-TanggalInput")
	// list := []TukarJaga{}
	// _, err = q.GetAll(ctx, &list)
	// if err != nil {
	// 	ErrorRec(ctx, "Gagal mengambil data", err)
	// 	SendBackError(w, "Gagal mengambil data", 500)
	// 	return
	// }
	// log.Infof(ctx, "Isi list adalah: %v", list)
	SendBackSuccess(w, nil, GenTemplate(ctx, list, "hal-tukar-jaga-content"), "Berhasil menambahkan data", "")
	// log.Infof(ctx, "Berhasil menambahkan data")

}
func getListTukarJaga(c context.Context) ([]TukarJaga, error) {
	q := datastore.NewQuery("TukarJaga").Order("-TanggalInput")
	list := []TukarJaga{}
	t := q.Run(c)
	for {
		tu := &TukarJaga{}
		k, err := t.Next(tu)
		if err == datastore.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		if tu.Status == "4" {
			continue
		}
		tu.Link = k.Encode()
		list = append(list, *tu)
	}
	// k, err := q.GetAll(c, &list)
	// if err != nil {
	// 	// ErrorRec(ctx, "Gagal mengambil data", err)
	// 	// SendBackError(w, "Gagal mengambil data", 500)
	// 	return nil, err
	// }
	// for m, n := range list {
	// 	n.Link = k[m].Encode()
	// }
	return list, nil

}
func getTukarJaga(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	list, err := getListTukarJaga(ctx)
	if err != nil {
		ErrorRec(ctx, "Gagal mengambil list tukar jaga", err)
		SendBackError(w, "Gagal mengambil list tukar jaga", 500)
	}
	// q := datastore.NewQuery("TukarJaga").Order("-TanggalInput")
	// list := []TukarJaga{}
	// _, err := q.GetAll(ctx, &list)
	// if err != nil {
	// 	ErrorRec(ctx, "Gagal mengambil data", err)
	// 	SendBackError(w, "Gagal mengambil data", 500)
	// 	return
	// }
	// log.Infof(ctx, "List adalah: %v", list)
	w.WriteHeader(200)
	fmt.Fprint(w, GenTemplate(ctx, list, "hal-tukar-jaga-content"))
	// SendBackSuccess(w, nil, GenTemplate(ctx, list, "hal-tukar-jaga-content"), "", "")

}

func ubahTglBayarJaga(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	js := &CatchDataJson{}
	json.NewDecoder(r.Body).Decode(js)
	defer r.Body.Close()
	// log.Infof(ctx, "isi dari body adalah: %v", js.Data3)
	k, err := datastore.DecodeKey(js.Data3)
	if err != nil {
		SendBackError(w, "Gagal mendecode key", 500)
		return
	}
	jag := &TukarJaga{}
	err = datastore.Get(ctx, k, jag)
	log.Infof(ctx, "Kunci adalah: %v", k)
	if err != nil {
		ErrorRec(ctx, "Gagal mengambil data", err)
		SendBackError(w, "Gagal mengambil data ", 500)
		return
	}
	jag.TanggalJagaBayar = ChangeStringtoTime(js.Data2)
	jag.JagaBayar = js.Data1
	_, err = datastore.Put(ctx, k, jag)
	if err != nil {
		SendBackError(w, "Gagal menyimpan data", 500)
		return
	}
	list, err := getListTukarJaga(ctx)
	if err != nil {
		SendBackError(w, "Gagal mengambil data", 500)
		return
	}
	// log.Infof(ctx, "Hal baru : %v", GenTemplate(ctx, list, "hal-tukar-jaga-content"))
	SendBackSuccess(w, nil, GenTemplate(ctx, list, "hal-tukar-jaga-content"), "Berhasil mengubah data", "")
	// GenTemplate(ctx, list, "hal-tukar-jaga-content")
}

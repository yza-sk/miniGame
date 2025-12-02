package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"miniGame/internal/svc"
)

type submitReq struct {
	Name    string `json:"name"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

type submitResp struct {
	Ok bool `json:"ok"`
}

type recordVO struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Score    int    `json:"score"`
	Comment  string `json:"comment"`
	Finished string `json:"finished"`
}

func SubmitHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req submitReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid body")
			return
		}
		if req.Name == "" || req.Score < 0 {
			writeErr(w, http.StatusBadRequest, "name required & score >=0")
			return
		}
		if err := ctx.Record.Insert(req.Name, req.Score, req.Comment); err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, submitResp{Ok: true})
	}
}

func RankHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := parseLimit(r, 50)
		list, err := ctx.Record.Top(limit)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		resp := make([]recordVO, 0, len(list))
		for _, v := range list {
			resp = append(resp, recordVO{ID: v.ID, Name: v.Name, Score: v.Score, Comment: v.Comment, Finished: v.Finished.Format(timeFormat)})
		}
		writeJSON(w, resp)
	}
}

func RecentHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := parseLimit(r, 50)
		list, err := ctx.Record.Recent(limit)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		resp := make([]recordVO, 0, len(list))
		for _, v := range list {
			resp = append(resp, recordVO{ID: v.ID, Name: v.Name, Score: v.Score, Comment: v.Comment, Finished: v.Finished.Format(timeFormat)})
		}
		writeJSON(w, resp)
	}
}

const timeFormat = "2006-01-02T15:04:05Z07:00"

func parseLimit(r *http.Request, def int) int {
	q := r.URL.Query().Get("limit")
	if q == "" {
		return def
	}
	v, err := strconv.Atoi(q)
	if err != nil || v <= 0 || v > 200 {
		return def
	}
	return v
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

package activityroute

import (
	"net/http"

	log "github.com/cjburchell/go-uatu"
	"github.com/cjburchell/ridemanager/api/routes/token"
	"github.com/cjburchell/ridemanager/common/service/data"
	"github.com/cjburchell/ridemanager/common/service/data/models"
	"github.com/gorilla/mux"
)

type validateWritable struct {
	service data.IService
	log     log.ILog
}

func (v validateWritable) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		user, err := token.GetUser(req, v.service)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// check user role
		if user.Role == models.AdminRole {
			next(w, req)
			return
		}

		// check activity owner
		vars := mux.Vars(req)
		activityId := models.ActivityId(vars["ActivityId"])

		activity, err := v.service.GetActivity(activityId)
		if err != nil {
			v.log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if activity.Owner.Id != user.Athlete.Id {
			activityId := models.AthleteId(vars["AthleteId"])
			if activityId != user.Athlete.Id {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		next(w, req)
	})
}

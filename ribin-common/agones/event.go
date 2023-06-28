package agones

import (
	"sync"

	coresdk "agones.dev/agones/pkg/sdk"
)

var once sync.Once

func WatchEvents() {
	gs := GetGameServer()
	lastAllocated := gs.ObjectMeta.Annotations["agones.dev/last-allocated"]

	err := agonesSDK.WatchGameServer(func(gs *coresdk.GameServer) {
		switch gs.Status.State {
		case "Allocated":
			la := gs.ObjectMeta.Annotations["agones.dev/last-allocated"]
			if lastAllocated != la {
				//首次allocated
				lastAllocated = la
			}
			once.Do(func() {
				labels := gs.GetObjectMeta().GetLabels()
				if labels != nil {
				}
			})
		}
	})
	if err != nil {
		panic(err)
	}
}

var StartAsyncServerFunc func(mapId string) = nil

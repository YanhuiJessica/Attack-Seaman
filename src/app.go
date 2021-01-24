package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zGina/Attack-Seaman/src/config"
	"github.com/zGina/Attack-Seaman/src/database"
	"github.com/zGina/Attack-Seaman/src/mode"
	"github.com/zGina/Attack-Seaman/src/model"
	"github.com/zGina/Attack-Seaman/src/router"
	"github.com/zGina/Attack-Seaman/src/runner"
)

var (
	// Version the version of TMA.
	Version = "unknown"
	// Commit the git commit hash of this version.
	Commit = "unknown"
	// BuildDate the date on which this binary was build.
	BuildDate = "unknown"
	// Mode the build mode
	Mode = mode.Dev
)

func main() {
	vInfo := &model.VersionInfo{Version: Version, Commit: Commit, BuildDate: BuildDate}
	mode.Set(Mode)

	fmt.Println("Starting TMA version", vInfo.Version+"@"+BuildDate)
	rand.Seed(time.Now().UnixNano())
	conf := config.Get()

	db, err := database.New(conf.Database.Connection, conf.Database.Dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	engine := router.Create(db, vInfo, conf)
	runner.Run(engine, conf)
}

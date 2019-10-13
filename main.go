package main

import (
	"fmt"

	"github.com/sellerin/gclustercore"
	"github.com/spf13/cobra"
)

func main() {
	var gitRepo string
	var revision string
	var simulationName string
	var data string
	var nbInjectors int
	var nbVu int
	var duration int
	var ramp int

	var cmdLaunch = &cobra.Command{
		Use:   "launch",
		Short: "Launch test",
		Long:  `launch is for starting a performance test with Gatling.`,
		Run: func(cmd *cobra.Command, args []string) {

			t := &gclustercore.TestConfiguration{
				GitRepo:        gitRepo,
				Revision:       revision,
				SimulationName: simulationName,
				Data:           data,
				NbInjectords:   int32(nbInjectors),
				NbVirtualUsers: int32(nbVu),
				Duration:       int64(duration),
				Ramp:           int64(ramp),
			}

			id := gclustercore.LaunchTest(t, gclustercore.NamespaceDev)

			fmt.Printf("Gatling test started. Test id: %s\n", id)
		},
	}

	var cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Delete test",
		Long:  `delete is for deleting all previous jobs.`,
		Run: func(cmd *cobra.Command, args []string) {
			gclustercore.DeleteJobs(gclustercore.NamespaceDev)
			fmt.Printf("Jobs are deleted\n")
		},
	}

	cmdLaunch.Flags().StringVarP(&gitRepo, "gitrepo", "g", "https://$(GIT_USERNAME):$(GIT_PASSWORD)@github.com/sellerin/gatling-simulation.git", "gitrepo with tests")
	cmdLaunch.Flags().StringVarP(&revision, "revision", "r", "master", "gitrepo branch")
	cmdLaunch.Flags().StringVarP(&simulationName, "simulation", "s", "c2gwebaws.C2gwebSimulation", "name of the simulation")
	cmdLaunch.Flags().StringVarP(&data, "data", "j", "{}", "json input data")
	cmdLaunch.Flags().IntVarP(&nbInjectors, "nbinjectors", "i", 2, "number of injectors")
	cmdLaunch.Flags().IntVarP(&nbVu, "nbvu", "v", 5, "number of virtual users")
	cmdLaunch.Flags().IntVarP(&duration, "duration", "d", 300, "duration of the test")
	cmdLaunch.Flags().IntVarP(&ramp, "ramp", "p", 10, "init ramp")

	var rootCmd = &cobra.Command{Use: "gclustercli"}
	rootCmd.AddCommand(cmdLaunch)
	rootCmd.AddCommand(cmdDelete)
	rootCmd.Execute()
}

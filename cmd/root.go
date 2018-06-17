package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/picosushi/meshi/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "meshi",
	Short: "meshi returns random meshi for given parameters",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// Currently, randomMeshi is root command.
	Run: randomMeshi,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.meshi.yaml)")

	rootCmd.PersistentFlags().Float64("lat", 35.690921, "Latitude for center location")
	rootCmd.PersistentFlags().Float64("lng", 139.700258, "Longtitude for center location")
	rootCmd.PersistentFlags().Int("distance", 500, "Distance from given latitude/longtitude")
	rootCmd.PersistentFlags().String("keyword", "いちごパフェ", "keyword for searching meshi")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".meshi" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".meshi")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func randomMeshi(cmd *cobra.Command, args []string) {
	// fmt.Println(args)
	// fmt.Println(cmd)

	lat, _ := cmd.Flags().GetFloat64("lat")
	lng, _ := cmd.Flags().GetFloat64("lng")
	distance, _ := cmd.Flags().GetInt("distance")
	keyword, _ := cmd.Flags().GetString("keyword")

	api_key := os.Getenv("GOOGLE_MAPS_API_KEY")

	response := meshi.Meshi(api_key,
		lat, lng,
		uint(distance),
		keyword,
	)

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	if response.NextPageToken != "" {
		fmt.Println(response.NextPageToken)
	}
}

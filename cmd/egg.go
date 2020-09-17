package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(eggCmd())
}

func eggCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "egg",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			printBUMP()
		},
	}
	cmd.SetHelpFunc(func(*cobra.Command, []string) {})
	cmd.SetUsageFunc(func(*cobra.Command) error { return nil})
	return cmd
}

func printBUMP() {
	fmt.Println("                                    `.-:/+oo+++++++o+/:.`                   ")
	fmt.Println("                                `./syo+//:----::://+osddy+:`                ")
	fmt.Println("                             `./so/+++ooo++++///:::----.../s+.              ")
	fmt.Println("                           `/sdy/:-......`.`````````````````/h/`            ")
	fmt.Println("                          -y+.--::/++oooooooooooooooooosssooosys`           ")
	fmt.Println("                         /h-                         ```````..-hs           ")
	fmt.Println("                       `+mo//::------::::::://+++oosssyyyyyhhhyym:          ")
	fmt.Println("                      `sdssssyyyyyyyyyyyyyyyyssssoooooooooooooooyd`         ")
	fmt.Println("                     `hhooooooooooooooooooooooooooooooooooooooooom:         ")
	fmt.Println("                     shooooooooooooooooydhoooosddoooooooooooooooohs         ")
	fmt.Println("             /yyo.  -moooooooooooooooosNNmoooomNNsoooooooooooooooyh         ")
	fmt.Println("            `mhoym-.shoooooooooooooooooddsoooohdyoooooooooooooooohy `:syo`  ")
	fmt.Println("          `shddyohhhmyooooooooooymooooooooooooooooooooooooooooooodh/hdosms  ")
	fmt.Println("          .hdyhyoooomsoooooooooohmoooooooooooooooooooooooooooooooNyyhoydms: ")
	fmt.Println("           `dmsshoshmhsoooooooooyNoooooooooooooooooysooooooooooohmoooohysdd`")
	fmt.Println("            -syNhom+dsshhysoooooohdooooooooooooooosmyooooooooosymhdohyshN+. ")
	fmt.Println("              .yhy+`+y-..:+syhhysshdysooooooooooosmhosssyhhhysyNo.moyNyy+`  ")
	fmt.Println("               ```  `ydsoo+++////+oyyhhhhhhhhhhhhddysso+/-.``-do` :yys``    ")
	fmt.Println("                     `y+``..-/+ooooooo++////::::://::::/+oosdd+`   ```      ")
	fmt.Println("                      `ss+:-.`  ```...-:://++++++//////:---sd:              ")
	fmt.Println("                       `/hdhyyyso+/:--..``              ./ho.               ")
	fmt.Println("`                        `:osyyssssyyyyyyyyssoo+++++osshyo.                 ")
	fmt.Println("                            `./oyyyssssooooosssssyhhhhs:.                   ")
	fmt.Println("`                               `.-://+odsoodyooomo:.`                      ")
	fmt.Println("`                                 `..---dhoomhooomo/::---..`                ")
	fmt.Println("`                             `-+syyyyyyysooddooosssyyyyyyyyo-              ")
	fmt.Println("`                            `hhysooooooossyddysssssooooooosdy              ")
	fmt.Println("`                            `syyyyyyyyyso+:--::/++oosssssso/. ``           ")
}


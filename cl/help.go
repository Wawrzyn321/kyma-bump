package cl

import (
	"bump/mappings"
	"bump/requirements"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"strings"
)

func PrintHelp() {
	printBranding()
	printCommandsHelp()
	requirements.PrintPathsRequirement()
	requirements.PrintDockerRequirement()
}

func printBranding() {
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println("* *                  Mr. BUMP                  * *")
	fmt.Println(strings.Repeat("*", 50))
}

func PrintBUMP() {
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

func printCommandsHelp() {
	fmt.Println("  Commands:")
	fmt.Println(Bold("    auto"))
	fmt.Println("      Diffs current state of Kyma and Console and bumps images. Usage:")
	fmt.Println("        bump auto -c <console tag> -k <kyma tag>")
	fmt.Println("      You can use either commit hash or PR tag. In former case, at least 8 characters of tag is required.")
	fmt.Println("      Add --no-verify or -f flag to disable image check.")
	fmt.Println(Bold("    check-files"))
	fmt.Println("      Checks if YAML configuration files exist and their tag variable paths match.")
	fmt.Println(Bold("    img"))
	fmt.Println("      Updates tags of images. Usage:")
	fmt.Println("        bump img <tag1> <...images> <tag2> <...images>")
	fmt.Println("      You can use either commit hash or PR tag. In former case, at least 8 characters of tag is required.")
	fmt.Println("      Add --no-verify or -f flag to disable image check.")
	fmt.Println(Bold("    help, -h"))
	fmt.Println("      Prints help.")
	fmt.Println(Bold("    list, -l"))
	fmt.Println("      Lists all supported images, along with their aliases.")
	fmt.Println(Bold("    verify"))
	fmt.Println("      Verifies changed images in repo.")
	fmt.Println("      You can pass branch to diff changes by. Defaults to HEAD.")
	fmt.Println("")
}

func List(m mappings.Mappings) {
	fmt.Println("Supported images:")
	for _, mapping := range m {
		fmt.Printf("%30s     %s\n", mapping.Name, strings.Join(mapping.Aliases, ", "))
	}
	fmt.Println("Yup, I dunno how to format in Go.")
}

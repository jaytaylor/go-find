package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jaytaylor/go-find"
)

var (
	MinDepth  int
	MaxDepth  int
	Type      string
	Name      string
	WholeName string
	Regex     string
	Empty     bool
	Print0    bool
	Mount     bool
)

var fs *flag.FlagSet

func init() {
	fs = flag.NewFlagSet("go-find", flag.ExitOnError)
	fs.IntVar(&MinDepth, "mindepth", math.MinInt, "Do not apply any tests or actions at levels less than levels (a non-negative integer).  -mindepth 1 means process all files except the starting-points.")
	fs.IntVar(&MaxDepth, "maxdepth", math.MinInt, "Descend at most levels (a non-negative integer) levels of directories below the starting-points.  -maxdepth 0 means only apply the tests and actions to the starting-points themselves.")
	fs.StringVar(&Type, "type", "", `File is of type c:
			c      character (unbuffered) special
			d      directory
			p      named pipe (FIFO)
			f      regular file
			l      symbolic link; this is never true if the -L option or the -follow option is in effect, unless the symbolic link is broken.  If you want to search for symbolic links when -L is in effect, use -xtype.
			s      socket
		To search for more than one type at once, you can supply the combined list of type letters separated by a comma `+"`"+`,' (GNU extension).`)
	fs.StringVar(&Name, "name", "", `Base of file name (the path with the leading directories removed) matches shell pattern pattern.  Because the leading directories are removed, the file names considered for a match with -name will never include a slash, so `+"`"+`-name a/b' will never match anything (you probably need to use -path instead).`)
	fs.StringVar(&WholeName, "wholename", "", `File name matches shell glob pattern.`)
	fs.StringVar(&Regex, "regex", "", "File name matches regular expression pattern.  This is a match on the whole path, not a search.")
	fs.BoolVar(&Empty, "empty", false, "File is empty and is either a regular file or a directory.")
	fs.BoolVar(&Mount, "mount", false, "Restrict the search to the given path filesystem.")
	fs.BoolVar(&Print0, "print0", false, "Print the full file name on the standard output, followed by a null character (instead of the newline character).  This allows file names that contain newlines or other types of white space to be correctly interpreted by programs that process the find output.  This option corresponds to the -0 option of xargs.")

	fs.Usage = func() {
		fmt.Printf("Usage:\n\n\t%s [flags..] [paths...]\n", os.Args[0])
		fmt.Print("\n\tAvailable predicate tests:\n")
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\n\t-%v    \t%v\n", f.Name, f.Usage) // f.Name, f.Value
		})
	}
}

func getSearchPathsFromCommandLine(args []string) []string {
	var paths []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			break
		}
		paths = append(paths, filepath.Clean(arg))
	}
	return paths
}

func main() {
	args := os.Args[1:]
	paths := getSearchPathsFromCommandLine(args) // find the search paths to be traversed after the command name

	fs.Parse(args[len(paths):]) // parse the flags after command name and search paths
	if len(args) > 1 && (args[1] == "-h" || args[1] == "-help" || args[1] == "--help") {
		flag.Usage()
		return
	}
	if len(args) == 0 {
		// Default to CWD.
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current working directory: %s", err)
		}
		paths = append(paths, cwd)
	}

	finder := find.NewFind(paths...)
	if MinDepth != math.MinInt {
		finder = finder.MinDepth(MinDepth)
	}
	if MaxDepth != math.MinInt {
		finder = finder.MinDepth(MinDepth)
	}

	if Type != "" {
		finder = finder.Type(Type)
	}
	if Name != "" {
		finder = finder.Name(Name)
	}
	if WholeName != "" {
		finder = finder.WholeName(WholeName)
	}
	if Mount {
		finder = finder.Mount()
	}
	if Regex != "" {
		expr, err := regexp.Compile(Regex)
		if err != nil {
			log.Fatalf("invalid regular expression %q: %s", Regex, err)
		}
		finder = finder.Regex(expr)
	}
	if Empty {
		finder = finder.Empty()
	}
	results, err := finder.Evaluate()
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		if Print0 {
			fmt.Printf("%s%v", result, byte(0))
		} else {
			fmt.Println(result)
		}
	}

	//if err := rootCmd.Execute(); err != nil {
	//	log.Fatal(err)
	//}
}

//var rootCmd = &cobra.Command{
//	Use: "find",
//	//Short: "",
//	//Long:  "",
//	//Args: cobra.
//	PersistentPreRun: func(_ *cobra.Command, _ []string) {
//	},
//	Run: func(_ *cobra.Command, args []string) {
//		if len(args) == 0 {
//			// Default to CWD.
//			cwd, err := os.Getwd()
//			if err != nil {
//				log.Fatalf("Error getting current working directory: %s", err)
//			}
//			args = append(args, cwd)
//		}
//
//		finder := find.NewFind(args...)
//		if MinDepth != math.MinInt {
//			finder = finder.MinDepth(MinDepth)
//		}
//		if MaxDepth != math.MinInt {
//			finder = finder.MinDepth(MinDepth)
//		}
//		if Type != "" {
//			finder = finder.Type(Type)
//		}
//		if Name != "" {
//			finder = finder.Name(Name)
//		}
//		if WholeName != "" {
//			finder = finder.WholeName(WholeName)
//		}
//		if Regex != "" {
//			expr, err := regexp.Compile(Regex)
//			if err != nil {
//				log.Fatalf("invalid regular expression %q: %s", Regex, err)
//			}
//			finder = finder.Regex(expr)
//		}
//		if Empty {
//			finder = finder.Empty()
//		}
//		results, err := finder.Evaluate()
//		if err != nil {
//			log.Fatal(err)
//		}
//		for _, result := range results {
//			if Print0 {
//				fmt.Printf("%s%v", result, byte(0))
//			} else {
//				fmt.Println(result)
//			}
//		}
//	},
//}

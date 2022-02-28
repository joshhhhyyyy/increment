package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/getsentry/sentry-go"
)

func main() {
	var finaltag string
	key := os.Getenv("key")
	nfpm := flag.Bool("nfpm", false, "Use output version number to nfpm")
	flag.Parse()

	uuuuuuuuu := sentry.Init(sentry.ClientOptions{
		Dsn:              key,
		TracesSampleRate: 1.0,
	})
	if uuuuuuuuu != nil {
		log.Println("sentry Init err")
		panic(uuuuuuuuu)
	}

	lmao, asdf := exec.Command("git", "describe", "--abbrev=0", "--tags").Output()
	if asdf != nil {
		log.Println(lmao)
		sentry.CaptureMessage(string(lmao))
		panic(asdf)
	}

	getlatesttag := string(lmao)
	log.Println("the initial (latest) tag is: ", getlatesttag)

	// set default values
	dotposition1 := 2
	dotposition2 := 3

	firstdigitposition1 := 1
	firstdigitposition2 := 2

	lastdigitposition1 := 5
	lastdigitposition2 := 6

	seconddigitposition1 := 3
	seconddigitposition2 := 4

	isdotthere := getlatesttag[dotposition1:dotposition2]

	for isdotthere != "." {
		dotposition1 = dotposition1 + 1
		dotposition2 = dotposition2 + 1

		firstdigitposition2 = firstdigitposition2 + 1

		seconddigitposition1 = seconddigitposition1 + 1
		seconddigitposition2 = seconddigitposition2 + 1

		lastdigitposition1 = lastdigitposition1 + 1
		lastdigitposition2 = lastdigitposition2 + 1

		isdotthere = getlatesttag[dotposition1:dotposition2]
	}

	getfirstdigit := getlatesttag[firstdigitposition1:firstdigitposition2]
	getseconddigit := getlatesttag[seconddigitposition1:seconddigitposition2]
	getlastdigit := getlatesttag[lastdigitposition1:lastdigitposition2]

	// if the last digit is 9, eg. v0.0.9,
	if getlastdigit == "9" {
		// make the last digit 0 (and add one to the second digit later)
		getlastdigit = "0"
		if getseconddigit == "9" {
			// if the second digit is also 9, eg. v0.9.9
			// make the second digit 0 {
			getseconddigit = "0"
			// } add one to the first digit {
			firstdigitconvertstringtonumber, wkauhfsevuiejroefw := strconv.Atoi(getfirstdigit) // Convert string to int
			newfirstdigit := firstdigitconvertstringtonumber + 1                               // add one
			getfirstdigit = strconv.Itoa(newfirstdigit)                                        // Convert int to string as per variable type
			// } result: 1.0.0
			if wkauhfsevuiejroefw != nil {
				sentry.CaptureException(wkauhfsevuiejroefw)
				panic(wkauhfsevuiejroefw)
			}
		} else {
			// else if it is not 9, eg. v0.8.9
			// add one to the second digit {
			seconddigitconvertstringtonumber, ueworiyiou4783788 := strconv.Atoi(getseconddigit) // Convert string to int
			newseconddigit := seconddigitconvertstringtonumber + 1                              // add one
			getseconddigit = strconv.Itoa(newseconddigit)                                       // Convert int to string as per variable type
			// } result: v0.9.0
			if ueworiyiou4783788 != nil {
				sentry.CaptureException(ueworiyiou4783788)
				panic(ueworiyiou4783788)
			}
		}
	} else {
		// if last digit is not 9, increment the last digit by 1
		lastdigitconvertstringtonumber, sahdiahd := strconv.Atoi(getlastdigit)
		if sahdiahd != nil {
			sentry.CaptureException(sahdiahd)
			panic(sahdiahd)
		}
		incrementlastdigit := lastdigitconvertstringtonumber + 1
		getlastdigit = strconv.Itoa(incrementlastdigit)
	}

	almostfinaltag := strings.Join([]string{getfirstdigit, getseconddigit, getlastdigit}, ".")
	finaltag = strings.Join([]string{"v", almostfinaltag}, "")

	log.Println("the new tag is: ", finaltag)

	gitfetch, fetcherr := exec.Command("git", "fetch").Output()
	if fetcherr != nil {
		log.Println(string(gitfetch))
		sentry.CaptureMessage(string(gitfetch))
		log.Println("there was an error when performing git fetch")
		panic(fetcherr)
	}

	gitpull, pullerr := exec.Command("git", "pull").Output()
	if pullerr != nil {
		log.Println(string(gitpull))
		sentry.CaptureMessage(string(gitpull))
		log.Println("there was an error when performing git push")
		panic(pullerr)
	}

	if *nfpm {
		log.Println("nfpm build ENABLED")
		os.Setenv("VERSION", finaltag)
		nfpmbuild, nfpmerr := exec.Command("nfpm", "package", "-p", "deb").Output()
		if nfpmerr != nil {
			log.Println(string(nfpmbuild))
			sentry.CaptureMessage(string(nfpmbuild))
			log.Println("there was an error when performing nfpm build")
			panic(nfpmerr)
		}
	}

	gitadd, adderr := exec.Command("git", "add", ".").Output()
	if adderr != nil {
		log.Println(string(gitadd))
		sentry.CaptureMessage(string(gitadd))
		log.Println("there was an error when performing git push")
		panic(adderr)
	}

	gitcommit, commiterr := exec.Command("git", "commit", "-m", "🫣").Output()
	log.Println(string(gitcommit))
	if commiterr != nil {
		sentry.CaptureMessage(string(gitcommit))
		log.Println("there was an error when performing git push")
		panic(commiterr)
	}

	gittag, tagerr := exec.Command("git", "tag", "-a", finaltag, "-m", "its new release time!! ✨").Output()
	if tagerr != nil {
		log.Println(string(gittag))
		sentry.CaptureMessage(string(gittag))
		log.Println("there was an error when performing git tag")
		panic(tagerr)
	}

	gitpushtag, pushtagerr := exec.Command("git", "push", "origin", finaltag).Output()
	if pushtagerr != nil {
		log.Println(string(gitpushtag))
		sentry.CaptureMessage(string(gitpushtag))
		log.Println("there was an error when performing git push tag")
		panic(pushtagerr)
	}

	gitpushmain, pushmainerr := exec.Command("git", "push").Output()
	if pushmainerr != nil {
		log.Println(string(gitpushmain))
		sentry.CaptureMessage(string(gitpushmain))
		log.Println("there was an error when performing git push main")
		panic(pushmainerr)
	}
}

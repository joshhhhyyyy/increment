const { exec } = require("child_process");

function increment(initialtag) {
  // set default values
  let dotposition1 = 2;
  let dotposition2 = 3;

  let firstdigitposition1 = 1;
  let firstdigitposition2 = 2;

  let lastdigitposition1 = 5;
  let lastdigitposition2 = 6;

  let seconddigitposition1 = 3;
  let seconddigitposition2 = 4;

  let isdotthere = initialtag.substring(dotposition1, dotposition2);

  while (isdotthere != ".") {
    dotposition1 = dotposition1 + 1;
    dotposition2 = dotposition2 + 1;

    firstdigitposition2 = firstdigitposition2 + 1;

    seconddigitposition1 = seconddigitposition1 + 1;
    seconddigitposition2 = seconddigitposition2 + 1;

    lastdigitposition1 = lastdigitposition1 + 1;
    lastdigitposition2 = lastdigitposition2 + 1;

    isdotthere = initialtag.substring(dotposition1, dotposition2);
  }

  let getfirstdigit = initialtag.substring(
    firstdigitposition1,
    firstdigitposition2
  );
  let getseconddigit = initialtag.substring(
    seconddigitposition1,
    seconddigitposition2
  );
  let getlastdigit = initialtag.substring(
    lastdigitposition1,
    lastdigitposition2
  );

  // if the last digit is 9, eg. v0.0.9,
  if (getlastdigit == "9") {
    // make the last digit 0 (and add one to the second digit later)
    getlastdigit = "0";
    if (getseconddigit == "9") {
      // if the second digit is also 9, eg. v0.9.9
      // make the second digit 0 {
      getseconddigit = "0";
      // } add one to the first digit {
      let firstdigitconvertstringtonumber = parseInt(getfirstdigit); // Convert string to int
      let newfirstdigit = firstdigitconvertstringtonumber + 1; // add one
      getfirstdigit = newfirstdigit.toString(); // Convert int to string as per variable type
      // } result: 1.0.0
    } else {
      // else if it is not 9, eg. v0.8.9
      // add one to the second digit {
      let seconddigitconvertstringtonumber = parseInt(getseconddigit); // Convert string to int
      let newseconddigit = seconddigitconvertstringtonumber + 1; // add one
      getseconddigit = newseconddigit.toString(); // Convert int to string as per variable type
      // } result: v0.9.0
    }
  } else {
    // if last digit is not 9, increment the last digit by 1
    let lastdigitconvertstringtonumber = parseInt(getlastdigit);
    let incrementlastdigit = lastdigitconvertstringtonumber + 1;
    getlastdigit = incrementlastdigit.toString();
  }

  let almostfinaltag = getfirstdigit + "." + getseconddigit + "." + getlastdigit;
  let finaltag = "v" + almostfinaltag;
  return finaltag;
}

function main() {
  let key = process.env.key;
  let nfpm = false;
  let dontpushmain = false;

  let lmao = exec("git describe --abbrev=0 --tags", (err, stdout, stderr) => {
    if (err) {
      console.log(stdout);
      console.log("there was an error when performing git fetch");
      throw err;
    }
    console.log("the initial (latest) tag is: ", stdout);
    let finaltag = increment(stdout);
    console.log("the new tag is: ", finaltag);

    exec("git fetch", (err, stdout, stderr) => {
      if (err) {
        console.log(stdout);
        console.log("there was an error when performing git fetch");
        throw err;
      }

      exec("git pull", (err, stdout, stderr) => {
        if (err) {
          console.log(stdout);
          console.log("there was an error when performing git push");
          throw err;
        }

        if (nfpm) {
          console.log("nfpm build ENABLED");
          process.env.VERSION = finaltag;
          exec("nfpm package -p deb", (err, stdout, stderr) => {
            if (err) {
              console.log(stdout);
              console.log("there was an error when performing nfpm build");
              throw err;
            }
          });
        }

        exec("git add .", (err, stdout, stderr) => {
          if (err) {
            console.log(stdout);
            console.log("there was an error when performing git push");
            throw err;
          }

          exec("git status --porcelain", (err, stdout, stderr) => {
            if (err) {
              throw err;
            }
            if (stdout.length == 0 && !dontpushmain) {
              throw "there are no errors to commit!";
            }

            exec("git commit -m 🫣", (err, stdout, stderr) => {
              if (err) {
                console.log(stdout);
                console.log("there was an error when performing git push");
                throw err;
              }

              exec(
                "git tag -a " + finaltag + " -m its new release time!! ✨",
                (err, stdout, stderr) => {
                  if (err) {
                    console.log(stdout);
                    console.log("there was an error when performing git tag");
                    throw err;
                  }

                  exec("git push origin " + finaltag, (err, stdout, stderr) => {
                    if (err) {
                      console.log(stdout);
                      console.log(
                        "there was an error when performing git push tag"
                      );
                      throw err;
                    }

                    if (!dontpushmain) {
                      exec("git push", (err, stdout, stderr) => {
                        if (err) {
                          console.log(stdout);
                          console.log(
                            "there was an error when performing git push main"
                          );
                          throw err;
                        }
                      });
                    } else {
                      console.log("Option:", " ", "Dont push main selected. Will not push to main.");
                      console.log(stdout);
                    }
                  });
                }
              );
            });
          });
        });
      });
    });
  });
}

main();

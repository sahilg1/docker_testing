This application is used to apply stress on the containers and meant for testing in Booster. The app is written in Go programming language.

The program creates 10 million threads which calculate the value of Pi in order to put load on the CPU and RAM.

The application connects to the S3 bucket to check for the continuity of the program to stop it.
https://s3.amazonaws.com/sahgupta-booster-stress-test/s3_file.txt

The app can be stopped by editing the text file by typing Stop/stop and saving it.

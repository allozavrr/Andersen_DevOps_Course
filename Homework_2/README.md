![alt text](https://github.com/allozavrr/Screenshots/blob/main/The_Elder_Scrolls_III_-_Morrowind_-_Text_Logo.png "Welcome to Elder Scripts, Traveler!")


# Welcome to the Morrowind script, Traveler. 
It will help you find out the status of your network connections.

Before starting work, please make sure that the whois and netstat utility spells are installed on your computer. 

The Morr_script must be downloaded from this repository in any way convenient for you. You need to set up privileges to enable the script to read, write, and execute. Use the Linux spell **chmod 751** for this.

As a great magician, you need to know the process name or PID, as well as by what criteria you need to filter the results.

To start the magic, run the command in the terminal and just type: ./morr_cript.sh

# How it works?

The Morr_script consists of several parts:

*  **Sudo** launches your spell on behalf of root. Feel like a real Nerevarine!

*  **Netstat** shows displays network connections. 
*The following superpowers netstat in this script shows TCP and UDP connections, theirs numerical addresses, the PID and name of the program (you can know this, Nerevarine!) and all listening sockets.

* **Awk** creates a report directly in the script itself.

* **Cut** it is a spell with which you extract information about the output of the data in each line. *You can adjust the number of lines here, Traveler.

* **Sort** will help you sort the results found alphabetically. *Well, or according to the runic alphabet.

* **Uniq** displays several similar values on one line. In this script, it will print the number of times that line is repeated in the input and one space before each line.

* **Tail** shows you how many pages from the end of the Elder Scrolls you have left. Or just the trailing lines of the output (you can type them yourself).

* **Grep** will help you find the desired rune word, in this case - will show the matching parts of the output.

* **While** will help you complete the magic - command...

* and **Read** can read it!

* **Whois** helps filter out IP connections *(you can filter the ones you need).

* **Make magic for your fun!**





![alt text](https://github.com/allozavrr/Screenshots/blob/main/620-6204264_elder-scrolls-morrowind-elder-scrolls-v-skyrim-elder.png "Morrowind never ending!")

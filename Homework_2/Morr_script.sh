echo "Welcome to Morrowind, sleepyhead. Even yesterday's rain didn't wake you up..."
echo "Keep calm and make whois for fun!"

#Check your priveleges before starting, Fighter!

while true; do
  read -p "Run the program as root? The Imperial guards are waiting for you... Are you sure? (YES/NO)" yn
  case $yn in
  [Yy]*)
    run_as_admin="y"
    break
    ;;
  [Nn]*)
    run_as_admin="n"
    break
    ;;
  *) echo "For the sake of the Emperor, choose an answer (YES/NO)" ;;
  esac
done

#You need write the process name, Nerevarine!

read -p "Choose your fighter! What is your process name or PID? " process_name

#You need type lines to see the process name, Nerevarine!

process_name="${process_name,,}"
read -p "Write the required number of lines for output (0-99): " lines #you can write lines you need, Traveler!
reg='^[0-9]{1,2}$' 
while [[ ! $lines =~ $reg ]]; do
  echo "The request is incorrect, please drink your skuuma and try again (0-99): "
  read lines
done

#Check your connections before starting, Fighter!

while true; do
  read -p "Enable viewing additional connection states? (YES/NO)" yn #you can wiew your connections, Traveler!
  case $yn in
  [Yy]*)
    add_states="y"
    break
    ;;
  [Nn]*)
    add_states="n"
    break
    ;;
  *) echo "Eat your moon sugar and choose an answer (YES/NO)" ;;
  esac
done

#Check standart whois filter spell, Fighter!

whois_filter="^Organization" #your guild, Traveler!

#Check your states of ports before starting, Fighter!

function show_add_info {
  echo "Traveler, you can find other connection states in the Elder Scroll below:" #you can wiew other connections, Traveler!
  netstat -tunapl | awk '/'$process_name'/ ''{print $5 "             " $6}'
  echo -e "_________________________________________________________________\n"
}

#Show your whois before starting, Fighter!
#you can wiew it in 5 lines... but as M'Aik the Liar says, this is not accurate

function show_whois {
  if [[ "$run_as_admin" == "y" ]]; then
    echo "Show your documents, traveler. Your password?"
    sudo netstat -tunapF | awk '/'$process_name'/ ''{print $5}' |
      cut -d: -f1 | sort | uniq -c | sort | tail -n"$lines" | grep -oP '(\d+\.){3}\d+' | while read IP; do whois $IP |
        awk -F':' '/'$whois_filter'/ ''{print $2}'; done | uniq -c
  else
    netstat -tunapF | awk '/'$process_name'/ ''{print $5}' |
      cut -d: -f1 | sort | uniq -c | sort | tail -n"$lines" | grep -oP '(\d+\.){3}\d+' | while read IP; do whois $IP | 
        awk -F':' '/'$whois_filter'/ ''{print $2}'; done | uniq -c
  fi
}

if [[ "$add_states" == "y" ]]; then
  show_add_info
  read -p 'Press [Enter] to continue and be patient, friend...'
  show_whois
else
  show_whois
fi

#Make your whois by the different filter, Fighter!

while true; do
  read -p "Would you like to change guild, traveler, in order to process the same whois request, but by the different filter? (Y/N) " yn
  case $yn in
  [Yy]*)
    run_again="y"
    break
    ;;
  [Nn]*)
    run_again="n"
    break
    ;;
  *) echo "Make your choice, Nerevarine! (YES/NO)" ;;
  esac
done

if [[ "$run_again" == "y" ]]; then
  read -p "Create a spell! Let new whois filter: " whois_filter
  show_whois 
fi 

# Completed the Dwemer Ruins of Dagoth Ur. So this was pretty confusing. But for you it was not difficult.
echo "Morrowind thanks you, Nerevarine!"


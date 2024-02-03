package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//example1()
	example2() 
 }

 func example1() {
	destinations := []string{
		"Relaxing on the golden sands in Greece",
		"Watching the northern lights in Finland",
		"Enjoying the Opera in Italy",
		"Shopping and sunbathing in Malaysia",
		"Drinking latte in a Cafe in Paris",
		"Experiencing flamenco in Spain",
		"Trekking in the Himalayas in Nepal",
		"Attending traditional tea ceremony in Japan",
	 }
	
	 fmt.Println("Welcome to the best ever, first time ever theeee TRAVEL GAME")
	 fmt.Println("Enter a number to find our your next travel destination and activity")
	
	 var input int
	
	 fmt.Println("Enter your number.")
	 fmt.Scan(&input)
	
	 src := rand.NewSource(time.Now().UnixNano() + int64(input))
	 r := rand.New(src)
	
	 //Randomly select a destination using the local random generator
	
	 destinationIndex := r.Intn(len(destinations))
	
	 chosenDestination := destinations[destinationIndex]
	
	 //Output the selected destination
	
	 fmt.Printf("Your travel adventure is: %s\n", chosenDestination)
 
	 //Provide specific advice based on the chosen destination
 
	 switch chosenDestination {
 
	 case "Relaxing on the golden sands in Greece":
		fmt.Println("Tip: Don't forget to wear sunscreen and visit the Islands!")
	 case "Watching the northern lights in Finland":
		fmt.Println("Tip: Make sure to bring a good camera to capture the aurora")
	 case "Enjoying the Opera in Italy":
		fmt.Println("Tip: Also try the gelato, Italy is famous for it!")
	 case "Shopping and sunbathing in Malaysia":
		fmt.Println("Tip: Bring cash hahahha and enjoy theme parks")
	 case "Drinking latte in a Cafe in Paris":
		fmt.Println("Explore the art galleries!")
	 case "Experiencing flamenco in Spain":
		fmt.Println("Indulge in the local cuisine, tapas")
	 case "Trekking in the Himalayas in Nepal":
		fmt.Println("Tip: Bring water")
	 case "Attending traditional tea ceremony in Japan":
		fmt.Println("Tip: Visit in cherry blossom season")
	
	 }
 }



func example2() {
 stages := []string{
  "early planning (1-6 months in)",
  "midway through planning (7-12 months in)",
  "final stages of planning (less than a month away)",
 }

	activities := [][]string{
		{
		 "Choosing your wedding theme and colors",
		 "Tasting sessions with potential caterers",
		 "Venue scouting for the ceremony and reception",
		},
		{
		 "DIY wedding invitation design workshop",
		 "Wedding dress and suit fitting day",
		 "Playlist creation for ceremony and reception",
		},
		{
		 "DIY centerpiece and decoration making session",
		 "Final walkthrough of the wedding day timeline",
		 "Confirming final details with all vendors",
		},
	 }

	 tips := [][]string{
		{
		 "Create a mood board to visualize your theme and color palette.",
		 "Rate each dish based on flavor, presentation, and uniqueness.",
		 "Take photos and notes on each venue's pros and cons for comparison.",
		},
		{
		 "Gather inspiration and use design software or templates for a personal touch.",
		 "Bring a trusted friend or family member for honest feedback and support.",
		 "Include a mix of genres and eras to cater to all guests' musical tastes.",
		},
		{
		 "Collect materials from craft stores or nature for unique, personal touches.",
		 "Discuss each moment with your partner and vendors to ensure smooth execution.",
		 "Review contracts and confirmations to avoid last-minute surprises.",
		},
	 }


	 fmt.Println("Welcome to the Wedding Planner's Ultimate Preparation Game!")
 fmt.Println("To get started, please tell us how far along you are in your wedding planning journey:")
 for i, stage := range stages {
  fmt.Printf("%d - %s\n", i+1, stage)
 }

 var stageInput int
 fmt.Println("Enter the number corresponding to your stage (or 0 to exit):")
 fmt.Scan(&stageInput)

 if stageInput == 0 {
  fmt.Println("Exiting the Wedding Planner's Game. Happy planning!")
  return
 } else if stageInput < 1 || stageInput > len(stages) {
  fmt.Println("Invalid input. Please restart the game and select a valid option.")
  return
 }

 stageIndex := stageInput - 1

 fmt.Printf("Based on your selection, here are activities suitable for %s:\n", stages[stageIndex])

 src := rand.NewSource(time.Now().UnixNano())
 r := rand.New(src)

 activityIndex := r.Intn(len(activities[stageIndex]))
 chosenActivity := activities[stageIndex][activityIndex]

 fmt.Printf("Your recommended activity is: %s.\n", chosenActivity)
 fmt.Println("Professional Tip:", tips[stageIndex][activityIndex])

 fmt.Println("Thank you for using the Wedding Planner's Game! Wishing you a beautiful journey to your big day.")
}

 
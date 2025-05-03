import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("Welcome to the text based dungeon game! Please enter your name: ");
        String playerName = scanner.nextLine();
        System.out.print("Hello, " + playerName + "! Are you ready to begin your journey? Yes or No?: ");
        String userInput = scanner.nextLine().toLowerCase();

       switch (userInput) {
           case "yes":
               System.out.println("Excellent! Let's begin!\n");
               break;

           case "no":
               System.out.println("Too bad! Game over.");
               scanner.close();
               return;
           default:
               System.out.println("Invalid input. Please enter 'Yes' or 'No'.");
               return;
       }
    // Rest of the game code goes here...
       System.out.print("You're in a dark and creepy hallway. As you walk down the hallway, you see a fork in it.\nOn the left, you see a door with dim light being cast under the door.\nTo the right, you see a door with bright light shining under it. Which will you choose? Left or Right?: ");
       String choice = scanner.nextLine().toLowerCase();

       switch (choice) {
           case "left":
               System.out.println("\nYou walk through the dimly lit door.\nYou can see that there is another door that is locked with a padlock on the opposite side of the room.\nMaybe we should check the desk in the corner for a key?\nYes or No?\n");
               String choice2 = scanner.nextLine().toLowerCase();
                switch (choice2) {
                     case "yes":
                         System.out.println("\nYou find the key and use it to open the padlock on the other door.\nIt leads to the outside.\nAs you step outside you hear the screams of the dumb ass that decided to go through the door on the right.\nCongratulations, you've escaped " + playerName + "! Great job you win!");
                         break;
                     case "no":
                         System.out.println("You decide not to search the desk. At that very moment the door behind you opens and it's a masked man with a machete.\nCongrats you done goofed, R.I.P. " + playerName + ". You played yourself.");
                         break;
                     default:
                         System.out.println("Invalid input. Please enter 'Yes' or 'No'.");
                         scanner.close();
                         break;
                  }

                     break;
           case "right":
               System.out.print("\nYou walk through the brightly lit door.\nAs you enter the room, you notice that there is an operating table in the middle.\nI think you know what comes next. R.I.P. " + playerName + ".");
               break;
           default:
               System.out.println("Invalid input. Please enter 'Yes' or 'No'.");
               scanner.close();
               break;
       }
       /*System.out.print("Would you like to search the desk in the corner? Yes or No?: ");
       choice = scanner.nextLine().toLowerCase();
       switch (choice) {
           case "yes":
               System.out.println("\nYou search the desk and find a small key. You use the key to open the padlock on the other door.\nIt leads to the outside.\nAs you step outside you hear the screams of the dumb ass that decided to go through the door on the right.\nCongratulations, you've escaped " + playerName + "! Great job you win!");
               break;
           case "no":
               System.out.println("You decide not to search the desk. At that very moment the door behind you opens and it's a masked man with a machete.\nCongrats you done goofed, R.I.P. " + playerName + ". You played yourself.");
               break;
           default:
               System.out.println("Invalid input. Please enter 'Yes' or 'No'.");
               scanner.close();
               break;*/
       }


        }



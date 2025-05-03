import java.util.Scanner;

public class Main {
  public static void main(String[] args) {
    int score1 = 1500;
    int score2 = 1000;
    int score3 = 500;
    int score4 = 100;
    int score5 = 50;
    String name = "Tim";


    calculateHighScore(score1);

    displayHighScore(1, name);

  }


  public static int displayHighScore(int position, String name) {

    if (position == 1) {
      System.out.println("Tim managed to get into position 1 on the high score list!");
    } else if (position == 2) {
      System.out.println("Tim managed to get into position 2 on the high score list!");
    } else if (position == 3) {
      System.out.println("Tim managed to get into position 3 on the high score list!");
    } else if (position == 4) {
      System.out.println("Tim managed to get into position 4 on the high score list!");
    }
    return position;
  }

  public static int calculateHighScore(int score){
    if (score >= 1000){
      System.out.println(1);
    }else if (score >= 500 && score < 1000){
      System.out.println(2);
    }else if (score >= 100 && score < 500){
      System.out.println(3);
    }else if (score > 100){
      System.out.println(4);
    }
    return score;
  }
}

// So I had some of the right ideas with bad execution. I didn't need to return anything for my first method.
// I'll make a new page showing the correct source code. I gave up too quickly trying to do the return code.
// Which is why I went with the println statements for the second method.
// I have a better understanding now though on how to make this work and tying a variable to a method.
// This way it can be passed.
public class MiniDuckSimulator {
 public static void main(String[] args) {
 Duck mallard = new MallardDuck();
 mallard.display();
 mallard.performFly();
 mallard.flyBehavior = new FlyNoWay();
 mallard.performFly();
 }
}
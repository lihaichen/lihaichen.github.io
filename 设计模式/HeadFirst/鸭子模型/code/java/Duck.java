
public abstract class Duck {
	FlyBehavior flyBehavior;
	public abstract void display();
	public void performFly() {
		flyBehavior.fly();
	}
}

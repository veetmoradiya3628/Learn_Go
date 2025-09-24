// Class
public class JedliksToyCar {
    private int batteryPercentage;
    private int distanceTravelled;
    
    public JedliksToyCar(){
        this.batteryPercentage = 100;
        this.distanceTravelled = 0;
    }
    
    public static JedliksToyCar buy() {
        return new JedliksToyCar();
    }

    public String distanceDisplay() {
        return "Driven " + this.distanceTravelled + " meters";
    }

    public String batteryDisplay() {
        if(this.batteryPercentage <= 0){
            return "Battery empty";
        }
        return "Battery at " + this.batteryPercentage + "%";
    }

    public void drive() {
       if (this.batteryPercentage > 0) {
            this.batteryPercentage -= 1;
            this.distanceTravelled += 20;
        }
    }
}

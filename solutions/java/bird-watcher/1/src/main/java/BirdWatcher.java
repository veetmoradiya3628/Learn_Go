
class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        int[] lastWeekCnt = {0, 2, 5, 3, 7, 8, 4};
        return lastWeekCnt;
    }

    public int getToday() {
        return this.birdsPerDay[this.birdsPerDay.length - 1];
    }

    public void incrementTodaysCount() {
        this.birdsPerDay[this.birdsPerDay.length - 1] += 1;
    }

    public boolean hasDayWithoutBirds() {
        for(int i = 0; i < this.birdsPerDay.length; i++){
            if(this.birdsPerDay[i] == 0) return true;
        }
        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        int cnt = 0;
        for(int i = 0; i < Math.min(numberOfDays, this.birdsPerDay.length); i++){
            cnt += this.birdsPerDay[i];
        }
        return cnt;
    }

    public int getBusyDays() {
        int cnt = 0;
        for(int i = 0; i < this.birdsPerDay.length; i++){
            if(this.birdsPerDay[i] >= 5) cnt++;
        }
        return cnt;
    }
}

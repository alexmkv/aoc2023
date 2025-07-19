package am.makeev;

public class Task0002
{
    public static long task02()
    {
        long v1 = 1;
        long v2 = 2;
        long sum = v2;
        while (true) {
            long v3 = v1 + v2;
            if (v3 > 4000000) break;
            v1 = v2;
            v2 = v3;
            if (v3 % 2 == 0) sum += v3;
        }
        return sum;
    }
    public static void main( String[] args )
    {
        System.out.println(task02());
    }
}
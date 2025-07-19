package am.makeev;

import java.util.stream.IntStream;

public class Task0001
{
    public static int task01()
    {
        return IntStream.range(0, 1000).filter(i -> i % 3 == 0 || i % 5 == 0).sum();
    }
    public static void main( String[] args )
    {
        System.out.println(task01());
    }
}

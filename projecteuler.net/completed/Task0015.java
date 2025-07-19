package am.makeev;

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashMap;
import java.util.Map;

public class Task015
{
    private static long lattice_paths(int n)
    {
        long[] grid = new long[n*n];
        Deque<Integer> toproc = new ArrayDeque<>();
        toproc.addLast(0);
        grid[0] = 1;
        while (!toproc.isEmpty()) {
            int v = toproc.pollFirst();
            int x = v/n;
            int y = v%n;
            if (x < n - 1) {
                int v1 = v + n;
                if (grid[v1] == 0)
                    toproc.addLast(v1);
                grid[v1] += grid[v];

            }
            if (y < n - 1) {
                int v1 = v + 1;
                if (grid[v1] == 0)
                    toproc.addLast(v1);
                grid[v1] += grid[v];
            }
        }
        return grid[(n-1)*n + n - 1];
    }

    public static void main( String[] args )
    {
        long startTime = System.currentTimeMillis();
        System.out.println(lattice_paths(2+1));
        System.out.println(lattice_paths(20));
        System.out.println(lattice_paths(20+1));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}
package am.makeev;

import java.util.HashMap;
import java.util.Map;

public class Task014
{
    private static int longest_collatz(int n)
    {
        int[] r = new int[n];
        long max = 0;
        int maxn = 0;
        for (int i = 1; i < n; ++i) {
            long j = i;
            int s = 0;
            while (j != 1) {
                if (j < n)
                {
                    int v = r[(int) j];
                    if (v != 0)
                    {
                        s += v;
                        break;
                    }
                }

                if (j % 2 == 0) j /= 2;
                else j = 3*j + 1;

                ++s;
            }
            r[i] = s;
            if (max < s)
            {
                max = s;
                maxn = i;
            }
        }
        return maxn;
    }

    public static void main( String[] args )
    {
        long startTime = System.currentTimeMillis();
        System.out.println(longest_collatz(1000000));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}
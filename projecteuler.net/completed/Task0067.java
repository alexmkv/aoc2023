package am.makeev;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.math.BigInteger;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Deque;
import java.util.List;
import java.util.stream.IntStream;
import java.util.stream.LongStream;

public class Task0067
{
    static int[][] data0 = {
            {3},
            {7, 4},
            {2, 4, 6},
            {8, 5, 9, 3}
    };


    private static int maximum_path_sum(int[][] data)
    {
        for (int i = 1; i < data.length; ++i) {
            int[] pd = data[i-1];
            int[] td = data[i];
            td[0] += pd[0];
            td[td.length-1] += pd[pd.length-1];
            for (int j = 1; j < td.length-1; ++j) {
                td[j] += Math.max(pd[j-1], pd[j]);
            }
        }
        return Arrays.stream(data[data.length-1]).max().orElse(0);
    }

    private static int[][] read_data(String filename) throws IOException
    {
        List<int[]> result = new ArrayList<>();
        BufferedReader br = new BufferedReader(new FileReader(filename));
        try {
            String line = br.readLine();

            while (line != null) {
                result.add(Arrays.stream(line.split(" ")).mapToInt(Integer::parseInt).toArray());
                line = br.readLine();
            }
        } finally {
            br.close();
        }
        int [][]ar = new int[result.size()][];
        return result.toArray(ar);
    }

    public static void main( String[] args ) throws IOException
    {
        long startTime = System.currentTimeMillis();
        System.out.println(maximum_path_sum(data0));
        //System.out.println(maximum_path_sum(read_data("Task0018.txt")));
        System.out.println(maximum_path_sum(read_data("Task0067.txt")));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}
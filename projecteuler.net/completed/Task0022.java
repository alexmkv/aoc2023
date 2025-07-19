package am.makeev;


import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.text.Collator;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Task0022
{

    private static long names_scores(List<String> names)
    {
        names.sort(Collator.getInstance());
        long res = 0;
        for (int i = 0; i < names.size(); ++i) {
            final long i1 = i+1;
            res += names.get(i).chars().mapToLong(v -> ((long)(v - 'A' + 1))*i1).sum();
        }
        return res;
    }

    private static List<String> read_data(String filename) throws IOException
    {
        List<String> result = new ArrayList<>();
        BufferedReader br = new BufferedReader(new FileReader(filename));
        try {
            String line = br.readLine();

            while (line != null) {
                Arrays.stream(line.split(",")).map(s -> s.substring(1, s.length()-1)).forEach(s -> result.add(s));
                line = br.readLine();
            }
        } finally {
            br.close();
        }
        return result;
    }

    public static void main( String[] args ) throws IOException
    {
        long startTime = System.currentTimeMillis();
        System.out.println(names_scores(read_data("Task0022.txt")));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}
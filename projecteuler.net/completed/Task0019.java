package am.makeev;

import java.util.Calendar;
import java.util.GregorianCalendar;


public class Task0019
{

    private static int count_mondays(Calendar startDate, Calendar endDate)
    {
        Calendar cur = startDate;
        int res = 0;
        int lastYear = 0;
        while(cur.get(Calendar.DAY_OF_WEEK) != Calendar.SUNDAY) {
            cur.add(Calendar.DAY_OF_MONTH, 1);
        }
        while(endDate.get(Calendar.DAY_OF_WEEK) != Calendar.SUNDAY) {
            endDate.add(Calendar.DAY_OF_MONTH, -1);
        }
        do {
            if (cur.get(Calendar.DAY_OF_MONTH) == 1)
                res++;
            cur.add(Calendar.DAY_OF_MONTH, 7);
            if (cur.get(Calendar.YEAR) != lastYear) {
                System.out.println(cur.get(Calendar.YEAR));
                lastYear = cur.get(Calendar.YEAR);
            }
        } while(cur.compareTo(endDate)<=0);
        return res;
    }


    public static void main( String[] args )
    {
        long startTime = System.currentTimeMillis();
        System.out.println(count_mondays(new GregorianCalendar(1901, 0, 1),
                new GregorianCalendar(2000, 11, 31)));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}
package am.makeev;

import java.math.BigInteger;
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.stream.IntStream;
import java.util.stream.LongStream;

public class Task0017
{
    private static int count_letters(int n)
    {
        switch (n) {
            case 1: return "one".length();
            case 2: return "two".length();
            case 3: return "three".length();
            case 4: return "four".length();
            case 5: return "five".length();
            case 6: return "six".length();
            case 7: return "seven".length();
            case 8: return "eight".length();
            case 9: return "nine".length();
            case 10: return "ten".length();
            case 11: return "eleven".length();
            case 12: return "twelve".length();
            case 13: return "thirteen".length();
            case 14: return "fourteen".length();
            case 15: return "fifteen".length();
            case 16: return "sixteen".length();
            case 17: return "seventeen".length();
            case 18: return "eighteen".length();
            case 19: return "nineteen".length();
            case 20: return "twenty".length();
            case 30: return "thirty".length();
            case 40: return "forty".length();
            case 50: return "fifty".length();
            case 60: return "sixty".length();
            case 70: return "seventy".length();
            case 80: return "eighty".length();
            case 90: return "ninety".length();
            case 1000: return count_letters(1) + "thousand".length();
        }
        if (n < 100) return count_letters((n / 10)*10) + count_letters(n % 10);
        if (n % 100 == 0) return count_letters(n / 100) + "hundred".length();
        return count_letters((n / 100) * 100) + "and".length()  + count_letters(n%100);
    }

    public static int count_all_letters(int n) {
        return IntStream.range(1, n + 1).map(v -> count_letters(v)).sum();
    }
/*
threehundredand
fortytwo

 */
    public static void main( String[] args )
    {
        long startTime = System.currentTimeMillis();
        System.out.println(count_all_letters(5));
        // 19954, 19964 - err
        System.out.println(count_letters(342)); // 23
        System.out.println(count_all_letters(1000));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}
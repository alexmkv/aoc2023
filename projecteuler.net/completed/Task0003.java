package am.makeev;

/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
 */
public class Task0003
{
    public static long largest_prime(long v)
    {
        for (int i = 2; i < v; ++i) {
            if (v % i == 0) {
                long v1 = largest_prime(i);
                long v2 = largest_prime(v/i);
                return v2 > v1 ? v2 : v1;
            }
        }
        return v;
    }
    public static void main( String[] args )
    {
        System.out.println(largest_prime(600851475143L));
    }
}
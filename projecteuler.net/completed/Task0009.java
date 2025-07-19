package am.makeev;




public class Task0009
{
    public static long triplet_product(int n)
    {
        for (int i0 = 1; i0 <= n - 2; ++i0) {
            for (int i1 = 1; i1 < n - i1; ++i1)
            {
                int i2 = n - i0 - i1;
                if (i2 <=0) continue;
                long productl = i1*i1 + i0*i0;
                long productr = i2*i2;
                if (productl > productr) break;
                if (productl == productr) return i1*i0*i2;
            }
        }
        return 0;
    }
    public static void main( String[] args )
    {
        System.out.println(triplet_product(12)); //60
        System.out.println(triplet_product(1000));
    }
}
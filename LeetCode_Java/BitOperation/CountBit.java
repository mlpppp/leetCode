package LeetCode_Java.BitOperation;

public class CountBit {
    public static void main(String[] args) {
        CountBit sol =  new CountBit();
        int bits = sol.countBit(12);
        System.out.println("num of bits in 12 is " + bits);
    }
    
    private int countBit(int num) {
        int count = 0;
        while (num > 0) {
            count += num & 1;
            num >>= 1;
        }
        return count;
    }
}

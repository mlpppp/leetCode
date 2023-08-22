package Interviews.TrendMicro;


class LongBit {
    public int solution(int A, int B) {
        // Implement your solution here
        long a = A;
        long b = B;
        long c = a*b;
        int count = 0;
        while(c > 0) {
            count += c&1;
            c >>= 1;
        }
        return count;
    }
    public static void main(String[] args) {
        int C = 100000000;
        // int C = 1024;
        // int C = 7;

        LongBit sol = new LongBit();
        System.out.println(Long.toBinaryString((long)C*C));
        System.out.println(sol.solution(C, C));
    }
}
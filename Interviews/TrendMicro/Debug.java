// you can also use imports, for example:
// import java.util.*;
package Interviews.TrendMicro;

// you can write to stdout for debugging purposes, e.g.
// System.out.println("this is a debug message");

class Debug {
    public boolean solution(int[] A, int K) {
        int n = A.length;
        for (int i = 0; i < n - 1; i++) {
            if (A[i] + 1 < A[i + 1])
                return false;
        }
        // vaild
        if (A[0] != 1 || A[n - 1] != K)
            return false;
        else
            return true;
    }

    public static void main(String[] args) {
        Debug aa = new Debug();
        int[] a = {0, 1, 1, 2, 3};
        System.out.println(aa.solution(a, 3));
    }
}

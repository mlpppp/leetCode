package Interviews.TrendMicro;


public class OneSwapSort {


    public boolean solution(int[] A) {
        // Implement your solution here
        if (A.length <= 1) return true;
        int firstOutsiderIdx = -1;
        int firstOutsiderLastOccurIdx = -1;

        // find first out of order number
        for (int i = 0; i < A.length-1; i++) {
            int j = i+1;
            if (A[j] < A[i]) {
                firstOutsiderIdx = j;
                break;
            }
        }

        // no out of order number, A is sorted
        if (firstOutsiderIdx < 0) return true;

        // find last occurance of the number
        for (int i = A.length-1; i >= 0; i--) {
            if (A[i] == A[firstOutsiderIdx]) {
                firstOutsiderLastOccurIdx = i;
                break;
            }
        }

        // swap the last occruance of outsider to the correct location (first location that larger than the outsider)
        for (int i = 0; i <= A.length-1; i++) {
            if (A[i] > A[firstOutsiderLastOccurIdx]) {
                swap(A, i, firstOutsiderLastOccurIdx);
                break;
            }
        }

        // verify if A is sorted
        return isSorted(A);
    }

    private void swap(int[] A, int i, int j) {
        if (i == j) return;
        int temp =  A[i];
        A[i] = A[j];
        A[j] =  temp;
        return;
    }

    private boolean isSorted(int[] A) {
        if (A.length <= 0) return true;
        for (int i = 0; i < A.length-1; i++) {
            int j = i+1;
            if (A[j] < A[i]) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        int[] A =  {5, 5};
        OneSwapSort sol = new OneSwapSort();
        System.out.println(sol.solution(A));
    }
}

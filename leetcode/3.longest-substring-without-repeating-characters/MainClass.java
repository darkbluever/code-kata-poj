import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.BitSet;
import java.util.Collections;
import java.util.List;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        BitSet bset = new BitSet(26);
        List list = new ArrayList();
        for (int i=0; i < s.length(); i++) {
            for(int j=i; j < s.length(); j++) {
                char ch = s.charAt(j);
                int ich = ch - ' ';

                if(bset.get(ich)) {
                    list.add(bset.cardinality());
                    bset.clear();
                    break;
                }

                bset.set(ich, true);
            }
        }

        if(!bset.isEmpty()) {
            list.add(bset.cardinality());
        }

        if (list.isEmpty()) {
            return 0;
        }

        return (int)Collections.max(list);
    }
}

class MainClass {
    public static String stringToString(String input) {
        if (input == null) {
            return "null";
        }
        return input.toString();
    }
    
    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            String s = stringToString(line);
            
            int ret = new Solution().lengthOfLongestSubstring(s);
            
            String out = String.valueOf(ret);
            
            System.out.print(out);
        }
    }
}

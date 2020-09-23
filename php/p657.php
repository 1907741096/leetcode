<?php
class Solution {

    /**
     * @param String $moves
     * @return Boolean
     */
    function judgeCircle($moves) {
        $hashSet = array(
            'R' => 1,
            'L' => -1,
            'U' => 1,
            'D' => -1,
        );
        $x = 0;
        $y = 0;
        for ($i = 0; $i < strlen($moves); $i ++) {
            if ($moves[$i] == 'R' || $moves[$i] == 'L') {
                $x += $hashSet[$moves[$i]];
            } else {
                $y += $hashSet[$moves[$i]];
            }
        }
        if ($x == 0 && $y == 0) {
            return true;
        }
        return false;
    }
}
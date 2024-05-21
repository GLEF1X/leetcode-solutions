MAX_ALLOWED_INTEGER = 2 ** 31 - 1
MIN_ALLOWED_INTEGER = -2 ** 31


class Solution:
    def myAtoi(self, s: str) -> int:
        s = s.lstrip()
        if s == "":
            return 0

        sign = 1
        possible_sign = s[0]
        if possible_sign == "+":
            s = s[1:]
        elif possible_sign == "-":
            sign = -1
            s = s[1:]
        
        if s == "":
            return 0

        is_checking_first_digits = True
        skipped_indexes = 0
        for i, char in enumerate(s):
            if not char.isdigit():
                s = s[:i - skipped_indexes]
                if s == "":
                    return 0

                return self.keep_integer_in_bounds_of_32_bits(int(s) * sign)

            if is_checking_first_digits and char == "0":
                s = s[1:]
                skipped_indexes += 1
                continue

            is_checking_first_digits = False

        if s == "":
            return 0

        return self.keep_integer_in_bounds_of_32_bits(int(s) * sign)

    @staticmethod
    def keep_integer_in_bounds_of_32_bits(i: int) -> int:
        return max(MIN_ALLOWED_INTEGER, min(i, MAX_ALLOWED_INTEGER))
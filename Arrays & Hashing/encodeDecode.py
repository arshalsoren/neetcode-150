# Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.
# Please implement encode and decode
# https://www.hackerrank.com/challenges/encode-and-decode/problem

class Solution:

    def encode(self, strs):
        res = ""
        for s in strs:
            res += str(len(s)) + "#" + s

        return res

    def decode(self, str):
        if str == "":
            return []

        res = []
        i = 0
        while i < len(str):
            j = str.find("#", i)
            if j == -1:
                break
            length = int(str[i:j])
            res.append(str[j+1:j+1+length])
            i = j+1+length

        return res

def main():
    strs = ["neet","code","love","you"]
    s = Solution()
    print("Encoded String: ", s.encode(strs))
    print("Decoded String: ", s.decode(s.encode(strs)))

if __name__ == '__main__':
    main()

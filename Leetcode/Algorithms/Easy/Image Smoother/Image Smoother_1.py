class Solution:
    def imageSmoother(self, img: List[List[int]]) -> List[List[int]]:
        newImg = [[img[i][j] for j in range(len(img[0]))] for i in range(len(img))]

        for i in range(len(img)):
            for j in range(len(img[0])):
                newImg[i][j] = self.average3x3(i, j, img)

        return newImg

    def average3x3(self, i: int, j: int, img: List[List[int]]) -> int:
        add, count = 0, 0
    
        for a in range(i - 1, i + 2):
            for b in range(j - 1, j + 2):
                if (a >= 0 and a < len(img)) and (b >= 0 and b < len(img[0])):
                    add += img[a][b]
                    count += 1

        return add // count
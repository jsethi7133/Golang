func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    R:=len(obstacleGrid)
    C:=len(obstacleGrid[0])

    if obstacleGrid[0][0]==1{
        return 0
    }

    obstacleGrid[0][0]=1

    for i:=1;i<R;i++{
        if (obstacleGrid[i][0]==0 && obstacleGrid[i-1][0]==1){
             obstacleGrid[i][0]=1
        }else{
            obstacleGrid[i][0]=0
        }

    }

    for i:=1;i<C;i++{
        if (obstacleGrid[0][i]==0 && obstacleGrid[0][i-1]==1){
            obstacleGrid[0][i]=1
        }else{
            obstacleGrid[0][i]=0
        }
    }

    for i:=1;i<R;i++{
        for j:=1;j<C;j++{
            if obstacleGrid[i][j]==0{
                obstacleGrid[i][j]=obstacleGrid[i-1][j]+obstacleGrid[i][j-1]
            }else{
                obstacleGrid[i][j]=0
            }
        }
    }
    return obstacleGrid[R-1][C-1]
}

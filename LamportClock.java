class LamportClock {
    
    static int max_ele(int a, int b) {
        if ( a > b) {
            return a;
        }
        else {
            return b;
        }
    }
    
    static void display(int e1, int e2, int p1[], int p2[]) {
        System.out.println("\n The order of events in process 1 ::");
        for(int i=0; i< e1; i++)
            System.out.print(p1[i] + " ");
            
        System.out.println("\n The order of events in process 2 ::");
        for(int j=0; j < e2; j++)
            System.out.print(p2[j] + " ");
    }
    
    static void lamportlogicalclock(int e1, int e2, int m[][]) {
        
        int p1[] = new int[e1];
        int p2[] = new int[e2];
        for(int i=0; i<e1; i++) {
            p1[i] = i + 1;
        }
        for(int j=0; j<e2;j++) {
            p2[j] = j + 1;
        }
        
        for(int j=0; j< e2; j++){
            System.out.print("\te2" + (j+1));
        }
        for(int i=0; i<e1; i++){
            System.out.print("\n e1" + (i+1) + "\t");
            for(int j=0; j<e2; j++) {
                 System.out.print(m[i][j] + "\t");
            }
           
        }
        
        for(int i=0;i<e1;i++){
            
            for(int j=0; j<e2; j++){
                if(m[i][j] == 1){
                    p2[j] = max_ele(p2[j], p1[i]+1);
                    for(int k=j+1;k<e2;k++)
                        p2[k] = p2[k-1] + 1;
                }
                
                if(m[i][j] == -1) {
                    p1[i] = max_ele(p1[i], p2[j]+1);
                    for(int k=i+1; k>e1;k++)
                        p1[k] = p1[k-1] + 1;
                }
            }
        }
        
        display(e1, e2, p1, p2);
    }
    
    public static void main(String[] args) {
        int events1 = 5, events2 = 3;
        int m[][] = new int[events1][events2];
        m[0][0] = 0;
        m[0][1] = 0;
        m[0][2] = 0;
        m[1][0] = 0;
        m[1][1] = 0;
        m[1][2] = 1;
        m[2][0] = 0;
        m[2][1] = 0;
        m[2][2] = 0;
        m[3][0] = 0;
        m[3][1] = 0;
        m[3][2] = 0;
        m[4][0] = 0;
        m[4][1] = -1;
        m[4][2] = 0;
        lamportlogicalclock(events1, events2, m);
    }
}


//         e21     e22     e23
//  e11    0       0       0
//  e12    0       0       1
//  e13    0       0       0
//  e14    0       0       0
//  e15    0       -1      0
//  The order of events in process 1 ::
// 1 2 3 4 5 
//  The order of events in process 2 ::
// 1 2 3 

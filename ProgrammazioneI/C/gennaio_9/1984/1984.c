//Autore: Francesco Corrado

#include <stdio.h>

typedef struct{
    int g, m, a;
} data;

conta_anno(data[]a, int n //Se non voglio resituire niente tipo void
        int anno){
    int i, c;
    c=0;
    for (i=0; i<n; i++)
        if (a[i].anno == anno)
            c++;
    return c
}

int main(){
    int n;
    printf("Quante date? ")
    scanf("%d", &n)
    data x[n];
    for (int i=0; i<n; i++){
        print("Giorno, mese, anno (separati da spazi)")
        scanf("%d%d%d", &x([i].giorno), &(x[i].mese), &(x[i].anno))
    }
    int conta;
    conta = conta - anno (x, n, 1987);
    
}
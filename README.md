# DSLR - Building Datascience tools & implementing Logistic Regression

*All scripts are compiled with **go build***

## Contents
The 42 DSLR project consists of 3 parts:
- **v1_data_analysis**<br>
  Building the Describe function from the well-known pandas library from scratch
- **v2_data_visualization**<br>
  Visualizing a dataset and it's relations using a histogram, scatter plot and pair plot
- **v3_logistic_regression**<br>
  Training a multiple classifier using 4 sigmoid classifiers

## V1: Data Analysis
### ./describe [csv file]
### Output example:
```
                Index           Feature1        Feature2        Feature3        
Count           4               4               4               4               
Mean            1.500000        283.393250      6751.162500     68.613750       
Std             1.290994        133.329037      3996.515534     24.812360       
Min             0.000000        137.643000      1231.730000     35.530000       
25%             0.250000        163.792250      2520.382500     43.760000       
50%             1.500000        269.345000      8010.340000     71.847500       
75%             2.750000        417.042500      9722.765000     90.233750       
Max             3.000000        457.240000      9752.240000     95.230000  
```

## V2: Data Visualization
### Histogram
v2_data_visualization/histogram/

### Scatter plot
v2_data_visualization/scatter_plot/

### Pair plot
v2_data_visualization/pair_plot/

## V3: Logistic Regression
### logreg_train
Train a model using ./logreg_train [dataset filename]<br>
- Generates a thetaValues.json file used by the prediction program
### logreg_predict
Predict using ./logreg_predict [dataset filename] [theta filename]
- Generates a houses.csv file with all predictions

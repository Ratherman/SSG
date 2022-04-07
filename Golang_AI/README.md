# Golang_AI
* ubuntu 20.04
* python: v3.7.0
* Anaconda: please get anaconda software from [here](https://www.anaconda.com/products/distribution)

```
conda create -n "Golang_AI" python=3.7.0

conda activate Golang_AI

conda install jupyter notebook

# this will spend more time
conda install -c conda-forge opencv 

# if only support cpu
conda install pytorch torchvision torchaudio cpuonly -c pytorch

# if support gpu
conda install pytorch torchvision torchaudio cudatoolkit=10.2 -c pytorch


conda install numpy

conda install matplotlib

conda install tqdm
```

# Classifier of Cat & Dog
* pytorch: v1.4.0 (cpu only)

# Reference
* [Datasets](https://www.kaggle.com/competitions/dogs-vs-cats/data?select=train.zip): kaggle dataset with 25,000 images in training dataset
* [Anaconda](https://www.anaconda.com/products/distribution)
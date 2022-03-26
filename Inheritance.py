from time import time
import matplotlib.pyplot as plt

class Model():

    """
    This Generic Model defines usual actions we will take 
    during the scope of Artificial Intelligence.

    Include: 

        - Dataset (Preprocessing), 
        - Train (AI Model), 
        - Evaluate (AI Model),
        - Save (AI Model)
    """

    def __init__(self, model_name, data_path, save_info, hparam):

        # Sanity Check
        assert type(model_name) == str
        assert type(data_path) == str
        assert type(save_info) == dict
        assert type(hparam) == dict

        # General
        self.model_name = model_name
        self.data_path = data_path
        self.save_info = save_info
        self.hparam = hparam

        # Statistic
        self.loss = []
        self.acc = []
        self.time = []

        print("\n\n=========================================================")
        print(f"Setup General and Statistic Parameters for {self.model_name}.")
        print("=========================================================\n\n")
        
    def dataset(self):
        """
        Display the location of dataset and pre-process.
        """

        print(f"Dataset is located at {self.data_path}")
        print(f"\n---- Preprocessing Start ----\n")

        return "\n---- Preprocessing Finished ---- \n"

    def train(self):
        """
        Display hyperparameters and start training
        """

        print(f"    epoch: {self.hparam['epoch']}")
        print(f"    learning rate: {self.hparam['lr']}")
        print(f"    batch size: {self.hparam['bs']}")
        print(f"\n---- Training Start ----\n")

        tic = time()
        for _ in range(self.hparam["epoch"]):
            
            # ------------------------------- #
            # 1. Forward Propagation          #
            # 2. Compute Loss                 #
            # 3. Backward Prorpagation        #
            # 4. Update Trainable Parameters  #
            # ------------------------------- #
            tmp_loss = 0
            tmp_acc = 0
            print(f"[{_}/{self.hparam['epoch']}], Loss: {'{:.4f}'.format(tmp_loss)}, Acc: {'{:.4f}'.format(tmp_acc)}")
            self.loss.append(tmp_loss)
            self.acc.append(tmp_acc)
        toc = time()
        self.time = round(toc - tic, 2)
        return "\n---- Training Finished ---- \n"

    def evaluate(self, plot_toggle=False):
        """
        Display information about training procedure.
        """
        print("\n---- Analysis Start ---- \n")
        print(f"Total Time: {self.time} (sec)")
        print(f"Loss: {round(self.loss[-1], 3)}") # get the last slot
        print(f"Accuracy: {round(self.acc[-1], 3)}") # get the last slot

        if plot_toggle == True:
            plt.figure(figsize=(10, 6))
            plt.title(f"Loss of {self.model_name}")
            plt.plot(self.loss)
            plt.show()

            plt.title(f"Accuracy of {self.model_name}")
            plt.plot(self.acc)
            plt.show()

        return "\n---- Analysis Finished ---- \n"

    def save(self):
        """Save the trained model."""

        print("\n---- Save Start ---- \n")
        print(f"{self.model_name} is saved as {self.save_info['save_name']}")
        print(f"{self.save_info['save_name']} is saved at {self.save_info['save_path']}")
        return "\n---- Save Finished ---- \n"

class MLP(Model):

    def __init__(self, save_name, save_path, epoch, lr, bs):
        self.model_name = "MLP"
        self.data_path = "./dataset/"
        self.save_info = {
            "save_name": save_name,
            "save_path": save_path
        }
        self.hparam = {
            "epoch": epoch,
            "lr": lr,
            "bs": bs,
        }

        super(MLP, self).__init__(
            self.model_name,
            self.data_path,
            self.save_info,
            self.hparam
        )

class ResNet(Model):

    def __init__(self, save_name, save_path, epoch, lr, bs):
        self.model_name = "ResNet"
        self.data_path = "./dataset/"
        self.save_info = {
            "save_name": save_name,
            "save_path": save_path
        }
        self.hparam = {
            "epoch": epoch,
            "lr": lr,
            "bs": bs,
        }

        super(ResNet, self).__init__(
            self.model_name,
            self.data_path,
            self.save_info,
            self.hparam
        )

class LSTM(Model):

    def __init__(self, save_name, save_path, epoch, lr, bs):
        self.model_name = "LSTM"
        self.data_path = "./dataset/"
        self.save_info = {
            "save_name": save_name,
            "save_path": save_path
        }
        self.hparam = {
            "epoch": epoch,
            "lr": lr,
            "bs": bs,
        }

        super(LSTM, self).__init__(
            self.model_name,
            self.data_path,
            self.save_info,
            self.hparam
        )

# =============== #
#  Neural Network #
# =============== #

JJs_NN = MLP(save_name="MLP_v1.h5", save_path="./trained_models/", epoch=30, lr=1e-3, bs=64)
print(JJs_NN.dataset())
print(JJs_NN.train())
print(JJs_NN.evaluate())
print(JJs_NN.save())

# ============================ #
# Convolutional Neural Network #
# ============================ #

JJs_ResNet = ResNet(save_name="ResNet_v1.h5", save_path="./trained_models/", epoch=30, lr=1e-3, bs=64)
print(JJs_ResNet.dataset())
print(JJs_ResNet.train())
print(JJs_ResNet.evaluate())
print(JJs_ResNet.save())

# ======================== #
# Recurrent Neural Network #
# ======================== #

JJs_LSTM = LSTM(save_name="LSTM_v1.h5", save_path="./trained_models/", epoch=30, lr=1e-3, bs=64)
print(JJs_LSTM.dataset())
print(JJs_LSTM.train())
print(JJs_LSTM.evaluate())
print(JJs_LSTM.save())
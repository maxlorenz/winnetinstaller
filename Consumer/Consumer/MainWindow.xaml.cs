using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace Consumer
{
    /// <summary>
    /// Interaktionslogik für MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {

        WebClient client;

        public MainWindow()
        {
            InitializeComponent();

            client = new WebClient();
            client.DownloadStringCompleted += new DownloadStringCompletedEventHandler(cmdExecuted);
        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {

            if (lstServers.SelectedIndex == -1)
            {
                return;
            }

            executeRequest(lstServers.SelectedItem.ToString(), txtCommand.Text);
        }

        private void cmdExecuted(object sender, DownloadStringCompletedEventArgs e)
        {
            if (e.Error == null)
            {
                txtOutput.Text = e.Result;
            }
            else
            {
                txtOutput.Text = e.Error.ToString();
            }
        }

        private void executeRequest(String server, String command)
        {

            try
            {
                client.DownloadStringAsync(new Uri("http://" + server + "/" + Uri.EscapeDataString(command)));
            }
            catch (Exception ex)
            {
                txtOutput.Text = ex.StackTrace;
            }
        }

        private void Button_Click_1(object sender, RoutedEventArgs e)
        {
            lstServers.Items.Add(txtServer.Text);
        }

        private void Button_Click_2(object sender, RoutedEventArgs e)
        {
            if (lstServers.SelectedIndex == -1)
            {
                return;
            }

            string cmdInfo = "wmic computersystem get domain, Manufacturer, Model, username";
            executeRequest(lstServers.SelectedItem.ToString(), cmdInfo);
        }

        private void Button_Click_3(object sender, RoutedEventArgs e)
        {
            if (lstServers.SelectedIndex == -1)
            {
                return;
            }

            executeRequest(lstServers.SelectedItem.ToString(), "test.bat");
        }
    }
}

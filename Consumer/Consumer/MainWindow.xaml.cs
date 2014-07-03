using System;
using System.Net;
using System.Windows;

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
        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {
            executeRequest(txtCommand.Text);
        }

        private void executeRequest(String command)
        {
            if (lstServers.SelectedIndex == -1)
            {
                return;
            }

            string server = lstServers.SelectedItem.ToString();
            webView.Navigate(new Uri("http://" + server + "/" + Uri.EscapeDataString(command)));
        }

        private void Button_Click_1(object sender, RoutedEventArgs e)
        {
            lstServers.Items.Add(txtServer.Text);
        }

        private void Button_Click_3(object sender, RoutedEventArgs e)
        {
            executeRequest("cmd /c echo list disk | diskpart");
        }

        private void Button_Click_4(object sender, RoutedEventArgs e)
        {
            executeRequest("\\\\wds01noham.bo.nexinto.com\\osdeploy\\webinterface\\install.bat");
        }

        private void Button_Click_5(object sender, RoutedEventArgs e)
        {
            executeRequest("wpeutil reboot");
        }

        private void Button_Click_6(object sender, RoutedEventArgs e)
        {
            if (lstServers.SelectedIndex == -1)
            {
                return;
            }

            infoView.Navigate(new Uri("http://" + lstServers.SelectedItem + "/_info"));
        }
    }
}

package com.example.tinyman

import android.os.Bundle
import android.util.Log
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import com.example.tinyman.databinding.FragmentExamplesBinding
import tinyman.Tinyman
import com.example.tinyman.getAccount
import com.example.tinyman.createClients
import com.example.tinyman.createTestAsset

/**
 * A simple [Fragment] subclass as the default destination in the navigation.
 */
class ExampleFragment : Fragment() {

    private var _binding: FragmentExamplesBinding? = null

    // This property is only valid between onCreateView and
    // onDestroyView.
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {

        _binding = FragmentExamplesBinding.inflate(inflater, container, false)
        return binding.root

    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        binding.textviewFirst.text = "Please choose an example. I suggest that you go to create an asset first"
        binding.buttonCreateAsset.setOnClickListener {
            findNavController().navigate(R.id.action_ExampleFragment_to_CreateAnAssetFragment)
        }
        binding.buttonBootstrap.setOnClickListener {
            findNavController().navigate(R.id.action_ExampleFragment_to_BootstrapFragment)
        }
        binding.buttonMint.setOnClickListener {
            findNavController().navigate(R.id.action_ExampleFragment_to_MintFragment)
        }
        binding.buttonBurn.setOnClickListener {
            findNavController().navigate(R.id.action_ExampleFragment_to_BurnFragment)
        }
        binding.buttonSwap.setOnClickListener {
            findNavController().navigate(R.id.action_ExampleFragment_to_SwapFragment)
        }
        binding.buttonRedeem.setOnClickListener {
            findNavController().navigate(R.id.action_ExampleFragment_to_RedeemFragment)
        }
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }
}
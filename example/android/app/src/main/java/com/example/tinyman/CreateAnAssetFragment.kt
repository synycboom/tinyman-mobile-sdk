package com.example.tinyman

import android.os.Bundle
import android.util.Log
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import com.example.tinyman.databinding.FragmentCreateAnAssetBinding

class CreateAnAssetFragment : Fragment() {

    private var _binding: FragmentCreateAnAssetBinding? = null

    // This property is only valid between onCreateView and
    // onDestroyView.
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {

        _binding = FragmentCreateAnAssetBinding.inflate(inflater, container, false)
        return binding.root

    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        binding.textviewOutput.text = "Click create to create an asset, it might take time to create"
        binding.buttonCreate.setOnClickListener {
            binding.buttonCreate.isEnabled = false
            binding.buttonCreate.isClickable = false
            binding.textviewOutput.text = "Loading..."
            SDKViewModel().createAsset {
                binding.buttonCreate.isEnabled = true
                binding.buttonCreate.isClickable = true
                Log.i("TINY_MAN_MOBILE_SDK", it)
                binding.textviewOutput.text = it
            }
        }
        binding.buttonGoBack.setOnClickListener {
            findNavController().navigate(R.id.action_CreateAnAssetFragment_to_ExampleFragment)
        }
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }
}